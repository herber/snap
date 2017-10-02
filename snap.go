package main

import (
	"fmt"
	"strings"
  "io/ioutil"
	"log"
  "os"
  "path/filepath"
	"strconv"
)

func main() {
	dir, err := os.Getwd()
  if err != nil {
      log.Fatal(err)
  }

	if len(os.Args) <= 1 {
		log.Fatal("Missing agrument")
	}

	if len(os.Args) >= 3 {
		log.Fatal("Too many agruments")
	}

  search := os.Args[1];
  if err != nil {
    log.Fatal(err)
  }

  if (search == "") {
    log.Fatal("No search string supplied");
  }

  check(dir, search, dir);
}

// test

func check(filename, search, dir string) {
  filename, _ = filepath.Abs(filename);

  if stat, err := os.Stat(filename); err == nil && stat.IsDir() {
    // "filename" is a directory
    files, err := ioutil.ReadDir(filename)
  	if err != nil {
  		log.Fatal(err)
  	}

  	for _, file := range files {
      f := filepath.Join(filename, file.Name())

  		check(f, search, dir)
  	}
  } else if _, err := os.Stat(filename); err == nil {
    // "filename" is a file
    content, err := ioutil.ReadFile(filename)
  	if err != nil {
  		log.Fatal(err)
  	}

    s := string(content[:])

    lines := strings.Split(s, "\n")

		print := false;
		basename := "." + filename[len(dir):]
		c := "\x1b[32m" + basename + "\x1b[0m\n";

    for i, line := range lines {
  		if strings.Contains(line, search) == true {
				i += 1;
				if (len(line) > 40) {
					c += "  \x1b[36m" + strconv.Itoa(i) + "\x1b[0m    " + line[:40] + "\n"
				} else {
					c += "  \x1b[36m" + strconv.Itoa(i) + "\x1b[0m    " + line + "\n"
				}

				print = true;
      }
   	}

		if print == true {
			fmt.Println(c)
		}
  }
}
