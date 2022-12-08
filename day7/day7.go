package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*************************************************
***               DISK object                  ***
*************************************************/

type Disk struct {
	PWD []*FileObject // a directory or file
	idx *FileObject   // ptr to current directory
}

func (d *Disk) backDir() {
	// do some pointer arithmetic to go back one directory
	if d.idx.name != "/" {
		// we are not at the root dir, so go back
		d.idx = d.idx.parent
	}
}

func (s *Disk) intoDir(dirName string) error {
	// first let's see if we need to initialize
	// fmt.Printf("Contents of os: +%v\n", s)
	if len(s.PWD) == 0 {
		// fmt.Println("Initializing")
		// we have to initialize.
		n := NewFileObject(dirName, 0, false)
		// it's assumed this is the root dir, so no parent
		s.PWD = append(s.PWD, &n)
		s.idx = s.PWD[len(s.PWD)-1]
		return nil
	}
	// first let's see if we have a record of this dir
	// already stated

	for i, v := range s.PWD {
		if dirName == v.name && v.parent == s.idx {
			// already have it, just update pointer
			s.idx = s.PWD[i]
			// fmt.Println("Using a known dir")
			return nil
		}
	}

	// otherwise, it's a new dir so let's initialize

	// fmt.Println("Adding new dir")
	a := NewFileObject(dirName, 0, false) // this is a dir so set accordingly
	a.AddParent(s.idx)
	// fmt.Printf("Contents of PWD before appending: +%v\n", s.PWD)
	s.PWD = append(s.PWD, &a)
	// fmt.Printf("Contents of PWD after appending: +%v\n", s.PWD)
	s.idx = s.PWD[len(s.PWD)-1]
	return nil
}

func (s *Disk) GetTotalDirSize() string {
	var retVal string = ""
	for _, v := range s.PWD {
		retVal += fmt.Sprintf("Total file size of %s: %d\n", v.name, v.Size())
	}
	return retVal
}

func AnySubDirs(f []*FileObject) bool {
	for _, v := range f {
		if !v.isFile {
			return false
		}
	}
	return true
}

func FindLessThan(f []*FileObject, totSize int) int {
	var totalSize int = 0
	for _, v := range f {
		// if v.parent == nil {
		// 	fmt.Printf("Dir: %s, Total size: %d\n", v.name, v.Size())
		// } else {
		// 	fmt.Printf("Dir: %s, Total size: %d, Parent: %s\n", v.name, v.Size(), v.parent.name)
		// }
		if v.Size() <= totSize {
			totalSize += v.Size()
		}
	}
	return totalSize
}

/*************************************************
***               FILE object                  ***
*************************************************/

type FileObject struct {
	name     string        // name of file or dir
	size     int           // size of file (0 if dir)
	isFile   bool          // true if a file, false if dir
	parent   *FileObject   // ptr to parent dir
	contents []*FileObject // if dir, will have ptrs to contents
}

func (f *FileObject) Size() int {
	// This is a recursive function which will return
	// the total size. If we are asking for the size
	// of a file, it will simply return the file size.
	// If it is a dir, it will recursively loop through
	// every file contents and add up the size.
	// fmt.Println("Working with:", f.name)
	// fmt.Printf("Is file? %t\n", f.isFile)

	if f.isFile {
		return f.size
	}
	totalSize := 0
	for _, v := range f.contents {
		totalSize += v.Size()
	}
	return totalSize
}

func NewFileObject(n string, s int, isf bool) FileObject {
	f := FileObject{}
	f.name = n
	f.size = s
	f.isFile = isf
	// separate function to insert parent
	return f
}

func (f *FileObject) AddParent(p *FileObject) {
	f.parent = p
}

func (f *FileObject) AddItem(c *FileObject) error {
	if f.isFile {
		return errors.New("Can only add contents to a directory type!")
	}

	f.contents = append(f.contents, c)
	return nil
}

func ParseLine(line string, o *Disk) {
	// this will just update the state
	// and OS file object
	out := strings.Split(line, " ")
	if out[0] == "$" {

		// this is a command.
		// either cd or ls
		if out[1] == "cd" {
			if out[2] == ".." {
				// fmt.Println("Going back...")
				o.backDir()
			} else {
				// fmt.Println("Going into", out[2])
				o.intoDir(out[2])
			}
		}
		// otherwise, ignore the ls i guess
	} else {
		if out[0] == "dir" {
			// First we need to see if this dir exists in
			// the main PWD slice, AND the parent is the same.
			checkParent := o.idx
			exists := false
			for _, v := range o.PWD {
				if v.name == out[1] && v.parent == checkParent {
					exists = true
					// already exists in main PWD, point to it
					// and move on
					o.idx.AddItem(v)
					break
				}
			}

			if !exists {
				// we looped through and couldn't find it,
				// time to create a new item
				n := NewFileObject(out[1], 0, false)
				n.AddParent(o.idx)
				o.idx.AddItem(&n)
				o.PWD = append(o.PWD, &n)
			}
		} else {
			// fmt.Println("Adding file", out[1])
			fileSize, err := strconv.Atoi(out[0])
			if err != nil {
				fmt.Println("Couldn't convert!", err)
			}
			n := NewFileObject(out[1], fileSize, true)
			n.AddParent(o.idx)
			o.idx.AddItem(&n)
		}
	}
}

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	thisOs := Disk{}

	for _, v := range lines {
		ParseLine(v, &thisOs)
	}

	fmt.Println("Total size of dirs under 100K:", FindLessThan(thisOs.PWD, 100000))

}
