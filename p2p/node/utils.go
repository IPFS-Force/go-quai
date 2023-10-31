package node

import (
	"bufio"
	"os"

	"github.com/dominant-strategies/go-quai/config"
	"github.com/dominant-strategies/go-quai/log"
	"github.com/spf13/viper"
)

// Utility function that asynchronously writes the provided "info" string to the node.info file.
// If the file doesn't exist, it creates it. Otherwise, it appends the new "info" as a new line.
func saveNodeInfo(info string) {
	go func() {
		// check if data directory exists. If not, create it
		dataDir := viper.GetString(config.DATA_DIR)
		if _, err := os.Stat(dataDir); os.IsNotExist(err) {
			err := os.MkdirAll(dataDir, 0755)
			if err != nil {
				log.Errorf("error creating data directory: %s", err)
				return
			}
		}
		nodeFile := dataDir + config.NODEINFO_FILE_NAME
		// Open file with O_APPEND flag to append data to the file or create the file if it doesn't exist.
		f, err := os.OpenFile(nodeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Errorf("error opening node info file: %s", err)
			return
		}
		defer f.Close()

		// Use bufio for efficient writing
		writer := bufio.NewWriter(f)
		defer writer.Flush()

		// Append new line and write to file
		log.Tracef("writing node info to file: %s", nodeFile)
		writer.WriteString(info + "\n")
	}()
}

// utility function used to delete any existing node info file
func deleteNodeInfoFile() error {
	dataDir := viper.GetString(config.DATA_DIR)
	nodeFile := dataDir + config.NODEINFO_FILE_NAME
	if _, err := os.Stat(nodeFile); !os.IsNotExist(err) {
		err := os.Remove(nodeFile)
		if err != nil {
			return err
		}
	}
	return nil
}
