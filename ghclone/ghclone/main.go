package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	git "gopkg.in/src-d/go-git.v4"
	"github.com/amclees/go-practice/ghclone"
)

func main() {
	var accessToken string
	var maxSizeKB int64 = 300 * 1000
	private := true
	ssh := true
	dest := "."

	switch len(os.Args) {
	case 1:
		fmt.Fprintln(os.Stderr, "No access token provided")
		printHelp()
		os.Exit(1)
	case 2:
		accessToken = os.Args[1]
	default:
		var help bool
		flag.StringVar(&accessToken, "access-token", "", "GitHub access token with repo scope, available at https://github.com/settings/tokens/new")
		flag.StringVar(&dest, "dest", "", "Destination where repositories will be cloned")
		flag.BoolVar(&private, "private", true, "Whether to include private repositories when cloning")
		flag.BoolVar(&ssh, "ssh", true, "Whether to use SSH links to download. HTTPS will be used when false.")
		flag.Int64Var(&maxSizeKB, "max-size", 300 * 1000, "Maximum size of repository in download, in kilobytes. Unlimited if less than 1, default 300 megabytes")
		flag.BoolVar(&help, "help", false, "Display help text")
		flag.Parse()

		if help {
			printHelp()
			return
		}
	}

	accessToken, ok := parseToken(accessToken)
	if !ok {
		fmt.Fprintln(os.Stderr, "That doesn't look like a GitHub access token")
		printHelp()
		os.Exit(1)
	}

	fmt.Println("Fetching repositories...")
	repos, err := ghclone.FetchRepos(accessToken, private)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching repos: %v\n", err)
		os.Exit(1)
	}

	existingFiles := make(map[string]bool)
	f, err := os.Open(dest)
	handleFileError(dest, err)
	filenames, err := f.Readdirnames(0)
	handleFileError(dest, err)
	for _, filename := range filenames {
		existingFiles[filename] = true
	}

	repoNames := make(map[string]bool)
	repoSkips := make(map[string]bool)
	for _, repo := range repos {
		if _, ok := repoNames[repo.Name]; ok {
			repoSkips[repo.Name] = true
		} else {
			repoNames[repo.Name] = true
		}
	}

	done := make([]chan bool, len(repos))
	for i := range done {
		done[i] = make(chan bool, 1)
	}
	for i, repo := range repos {
		if _, ok := existingFiles[repo.Name]; ok {
			fmt.Printf("File with name %s already exists; skipping...\n", repo.Name)
			done[i] <- true
			continue
		}
		if _, ok := repoSkips[repo.Name]; ok {
			fmt.Printf("Repo with full name %s collides with another repo's name; skipping...\n", repo.FullName)
			done[i] <- true
			continue
		}
		if maxSizeKB > 0 && repo.SizeKB > maxSizeKB {
			fmt.Printf("Repo with full name %s too big (%d kilobytes); skipping...\n", repo.FullName, repo.SizeKB)
			done[i] <- true
			continue
		}
		fmt.Printf("Cloning %s...\n", repo.Name)
		go fetch(dest, repo, ssh, done[i])
	}

	for _, ch := range done {
		_, _ = <-ch
	}
}

func fetch(dest string, repo ghclone.Repo, ssh bool, done chan bool) {
	url := repo.HTTPSURL
	if ssh {
		url = repo.SSHURL
	}

	_, err := git.PlainClone(path.Join(dest, repo.Name), false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error cloning repository %s: %v\n", repo.Name, err)
	}

	done <- true
}

func handleFileError(dest string, err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", dest, err)
	os.Exit(1)
}

func hex(r rune) bool {
	return (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F') || (r >= '0' && r <= '9')
}

func parseToken(str string) (string, bool) {
	s := make([]rune, 0, 0)
	for _, r := range str {
		if hex(r) {
			s = append(s, r)
		}
	}
	return string(s), len(s) == 40
}

func printHelp() {
	fmt.Fprintf(os.Stderr, `USAGE:
	ghclone [ACCESS TOKEN]
	ghclone [OPTION]...
	
Options:
	--access-token  GitHub access token
	--dest          Destination folder where repositories will be cloned (default .)
	--private       Whether to include private repositories when cloning (default true)
	--ssh           Whether to use SSH links to download. HTTPS will be used when false.` + "\n")
}

