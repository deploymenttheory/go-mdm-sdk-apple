package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-apple-sdk-devicemanagement/examples/helpers"
	"github.com/deploymenttheory/go-apple-sdk-devicemanagement/sdk/apple/mobiledevicemanagementprofiles/payloads/security/generate"

	"github.com/deploymenttheory/go-apple-sdk-devicemanagement/sdk/apple/serialize"
)

func main() {
	options := CreateSmartCardProfile()

	passcodeProfile := generate.CreateSmartCardMobileDeviceManagementProfile(options)

	// Serialize the profile to plist in XML format
	profileXML, err := serialize.SerializeDeviceManagementProfileToPLIST(passcodeProfile)
	if err != nil {
		fmt.Println("Error serializing profile to PLIST XML:", err)
		return
	}

	// Infer the filename from PayloadDisplayName and add the ".mobileconfig" suffix
	fileName := options.PayloadDisplayName + ".mobileconfig"

	// Rest of your file handling code...
	fmt.Print("Enter the directory to save the .mobileconfig file (e.g., /path/to/folder): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dirPath := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Join directory path and filename
	filePath := filepath.Join(dirPath, fileName)

	// Save the profile to the specified file path
	if err := helpers.SaveToFile(filePath, profileXML); err != nil {
		fmt.Println("Error saving profile to file:", err)
		return
	}
}

// CreateSmartCardProfile generates a profile with the provided options.
func CreateSmartCardProfile() generate.SmartCardMobileDeviceManagementProfileConfigOptions {
	return generate.SmartCardMobileDeviceManagementProfileConfigOptions{
		PayloadDisplayName:    "smartcard",
		UserPairing:           true,
		AllowSmartCard:        true,
		CheckCertificateTrust: 1,
		OneCardPerUser:        true,
		TokenRemovalAction:    1,
		EnforceSmartCard:      true,
	}
}
