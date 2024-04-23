cgrep command
=============

* [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Description
-----------

`cgrep` command that searches for words in files within a project (or wherever this command was executed from)

Features
--------

* **Finds exact matches**: Provide a string and will be looked in files inside of child folders.

Table of Contents
-----------------

* [Installation](#installation)
* [Usage](#usage)

### Installation

Recommended to create new folder directory (if you want the command as global)

```bash
mkdir ~/my-custom/path && cd ~/my-custom/path
```

Clone the repository:

```bash
git clone https://github.com/cjairm/cgrep.git

################################################################
### If you want to make it a global command...
echo 'alias cgrep="~/my-custom/path/cgrep/cgrep find"' >> ~/.zshrc
source ~/.zshrc
################################################################
```

### Usage
From here as simple as

```bash
cgrep "func("

// Response example:

// cmd/find.go
// 50:     visitAllFilesInDir(wd, func(pathToFile string) {
// 59: func visitAllFilesInDir(wd string, funcPathToFile func(pathToFile string)) {
// 60:     err := filepath.Walk(wd, func(path string, file fs.FileInfo, err error) error {

// cmd/root.go
// 19:     // Run: func(cmd *cobra.Command, args []string) { },
```

Future work
--------

* **Ignore files in .gitignore**
* **Search files of a specific type**

Enjoy! :smiley:
