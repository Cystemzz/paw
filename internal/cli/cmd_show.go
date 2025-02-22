package cli

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.design/x/clipboard"

	"lucor.dev/paw/internal/paw"
)

// Show shows an item details
type ShowCmd struct {
	itemPath
	clipboard bool
}

// Name returns the one word command name
func (cmd *ShowCmd) Name() string {
	return "show"
}

// Description returns the command description
func (cmd *ShowCmd) Description() string {
	return "Shows an item details"
}

// Usage displays the command usage
func (cmd *ShowCmd) Usage() {
	template := `Usage: paw-cli show [OPTION] VAULT_NAME/ITEM_TYPE/ITEM_NAME

{{ . }}

Options:
  -c, --clip  Do not print the password but instead copy to the clipboard
  -h, --help  Displays this help and exit
`
	printUsage(template, cmd.Description())
}

// Parse parses the arguments and set the usage for the command
func (cmd *ShowCmd) Parse(args []string) error {
	flags, err := newCommonFlags()
	if err != nil {
		return err
	}

	flagSet.BoolVar(&cmd.clipboard, "c", false, "")
	flagSet.BoolVar(&cmd.clipboard, "clip", false, "")

	flagSet.Parse(args)
	if flags.Help || len(flagSet.Args()) != 1 {
		cmd.Usage()
		os.Exit(0)
	}

	if cmd.clipboard {
		err := clipboard.Init()
		if err != nil {
			return err
		}
	}

	itemPath, err := parseItemPath(flagSet.Arg(0), itemPathOptions{fullPath: true})
	if err != nil {
		return err
	}
	cmd.itemPath = itemPath
	return nil
}

// Run runs the command
func (cmd *ShowCmd) Run(s paw.Storage) error {
	password, err := askPassword("Enter the vault password")
	if err != nil {
		return err
	}

	vault, err := s.LoadVault(cmd.vaultName, password)
	if err != nil {
		return err
	}

	item, err := paw.NewItem(cmd.itemName, cmd.itemType)
	if err != nil {
		return err
	}

	item, err = s.LoadItem(vault, item.GetMetadata())
	if err != nil {
		return err
	}

	var pclip []byte
	var pclipMsg string
	switch cmd.itemType {
	case paw.LoginItemType:
		v := item.(*paw.Login)
		fmt.Printf("URL: %s\n", v.URL)
		fmt.Printf("Username: %s\n", v.Username)
		if !cmd.clipboard {
			fmt.Printf("Password: %s\n", v.Password.Value)
		} else {
			pclip = []byte(v.Password.Value)
			pclipMsg = "[✓] password copied to clipboard"
		}
		if v.Note != nil {
			fmt.Printf("Note: %s\n", v.Note.Value)
		}
	case paw.PasswordItemType:
		v := item.(*paw.Password)
		if !cmd.clipboard {
			fmt.Printf("Password: %s\n", v.Value)
		} else {
			pclip = []byte(v.Value)
			pclipMsg = "[✓] password copied to clipboard"
		}
		if v.Note != nil {
			fmt.Printf("Note: %s\n", v.Note.Value)
		}
	case paw.SSHKeyItemType:
		v := item.(*paw.SSHKey)
		if !cmd.clipboard {
			fmt.Printf("Private key: %s\n", v.PrivateKey)
		} else {
			pclip = []byte(v.PrivateKey)
			pclipMsg = "[✓] private key copied to clipboard"
		}
		fmt.Printf("Public key: %s\n", v.PublicKey)
		fmt.Printf("Fingerprint: %s\n", v.Fingerprint)
		if v.Note != nil {
			fmt.Printf("Note: %s\n", v.Note.Value)
		}
	case paw.NoteItemType:
		v := item.(*paw.Note)
		fmt.Printf("Note: %s\n", v.Value)
	}

	fmt.Printf("Modified: %s\n", item.GetMetadata().Modified.Format(time.RFC1123))
	fmt.Printf("Created: %s\n", item.GetMetadata().Created.Format(time.RFC1123))

	if pclip != nil {
		ctx, cancel := context.WithTimeout(context.Background(), clipboardWriteTimeout)
		defer cancel()
		err := writeToClipboard(ctx, pclip)
		if err != nil {
			return nil
		}
		fmt.Println(pclipMsg)
	}
	return nil
}
