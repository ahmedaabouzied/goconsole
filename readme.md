# Go Console 

A golang implementation of some of the basic linux commands to interact with files and directories . 
I built this application to practice Go while I am learning about go libraries.

## Installation

    go install "github.com/ahmedaabouzied/goconsole"

## Usage 

Run ``` goconsole ``` if GOPATH is in your PATH.

Now the application prints out the current working directory and waits for your command .
If you run ``` ls ``` it prints out the current directory files and directories just like ``` ls ``` that is native in linux.

To exit the application just run ``` exit ```.

## Supported commands 

``` ls [directory name (optional)] ```

    list the contents of current directory.
    if a directory is provided it willlist that directory instead.

``` cd [directory name] ```

    Changes the current working directory into the directory provided.

``` touch [file name] ```

    Creates a new file with the provided file name.

``` del [file name] ```

    Deletes provided file from the current directory.

``` mkdir [directory name] ```

    Creates a new directory with the provided name.

``` copy [fileName  , newFileName]  ```
    
    Copies the file provided into the new file .


## TODO 

    * Support more commands 
    * Add tests

## Author

[Ahmed Abouzied](https://www.github.com/ahmedaabouzied)

Twitter [@ahmedaabouzied](https://twitter.com/ahmedaabouzied)

## Licence

MIT Licence , Do whatever the fuck you want with it.
