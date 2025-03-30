# Directory-Brusting-RFuzz
Directory and File Enumeration: The primary function. It discovers hidden directories and files on a web server by sending HTTP requests based on a provided wordlist

**Core Features**

Directory and File Enumeration: The primary function. It discovers hidden directories and files on a web server by sending HTTP requests based on a provided wordlist .
Color-Coded Output: Improves readability by displaying HTTP status codes in different colors:
200 OK: Green
301 Redirect: Yellow
403 Forbidden: Red
404 Not Found: Gray
Other status codes are displayed in a default color .
Concurrency (Fast Execution): Uses Go's goroutines to send multiple requests simultaneously, making the brute-forcing process faster .
Wordlist Support: Reads a wordlist file containing potential directory and file names .
Basic Usage: Simple command-line interface for specifying the target URL and wordlist .

**Advanced Features**

Customizable Threads: Allows users to specify the number of threads (concurrent requests) to use .
Output to File: Saves the results to a file for later analysis .
HTTP Header Customization: Enables users to add custom HTTP headers to the requests . Useful for bypassing certain security measures or simulating different user agents.
Proxy Support: Allows routing traffic through a proxy server .
Status Code Filtering: Lets users filter the output to only show results with specific HTTP status codes .
Timeout Configuration: Sets a timeout for HTTP requests .
Recursion: If a directory is found, Rfuzz can recursively scan that directory as well .
Wildcard Filters: Allows users to filter results based on wildcard patterns in the response .
User-Agent Spoofing: Randomly changes the User-Agent header in each request to avoid detection .
HTTPS Support: Properly handles HTTPS connections, including certificate verification .
Automatic Redirection Handling: Follows HTTP redirects .
Summary Table

**Feature	Description** 

Directory/File Enumeration	Discovers hidden web resources.
Color-Coded Output	Improves readability of results based on HTTP status codes.
Concurrency	Enables fast execution via multiple concurrent requests.
Wordlist Support	Reads potential directory/file names from a file.
Customizable Threads	Controls the number of concurrent requests.
Output to File	Saves results to a file.
HTTP Header Customization	Adds custom headers to HTTP requests.
Proxy Support	Routes traffic through a proxy server.
Status Code Filtering	Filters output based on HTTP status codes.
Timeout Configuration	Sets a timeout for HTTP requests.
Recursion	Scans subdirectories if found.
Wildcard Filters	Filters results based on patterns in the response.
User-Agent Spoofing	Randomizes the User-Agent header.
HTTPS Support	Handles secure HTTPS connections.
Automatic Redirection Handling	Follows HTTP redirects.
By incorporating these features, Rfuzz can become a versatile and powerful tool for web application security testing

**Installations Process:**

### Installation

1.  **Install Go:**

    *   First, ensure Go is installed on your Linux system. If not, download and install it from the official [Go website](https://go.dev/dl/).
    *   Verify the installation by running `go version` in the terminal.
2.  **Set up Go environment:**

    *   Configure the Go environment variables, such as `GOROOT` and `GOPATH`.
    *   Add `$GOPATH/bin` to your `PATH`.  You can typically do this by adding `export PATH=$PATH:$GOPATH/bin` to your `.bashrc` or `.zshrc` file.
3.  **Create the Rfuzz project:**

    *   Create a new directory for the Rfuzz project:

        ```bash
        mkdir Rfuzz
        cd Rfuzz
        ```
4.  **Write the Rfuzz code:**

    *   Create a file named `main.go` and add the Go code for Rfuzz.
5.  **Install dependencies:**

    *   Use `go get` to install any necessary dependencies.  For example:

        ```bash
        go get github.com/spf13/cobra
        ```
6.  **Build the Rfuzz binary:**

    *   Run `go build` to compile the `main.go` file into an executable binary:

        ```bash
        go build
        ```
7.  **Install Rfuzz:**

    *   Move the Rfuzz binary to a directory in your `PATH`, such as `/usr/local/bin`, to make it accessible system-wide:

        ```bash
        sudo mv Rfuzz /usr/local/bin
        ```

### Core Functionality

*   **Command-line arguments:** Use the `flag` package (or a more robust library like `cobra`) to handle command-line arguments such as the target URL, wordlist, number of threads, and output file [Reference 3]().
*   **HTTP requests:** Employ the `net/http` package to send HTTP requests to the target URL [Reference 3]().
*   **Concurrency:** Utilize Go's `goroutines` and `channels` to achieve concurrency, allowing for fast directory brute-forcing [Reference 10]().
*   **Color-coded output:** Use the `fmt` package along with ANSI escape codes to add color to the output based on HTTP status codes. For example:
    *   200 OK: Green
    *   301 Redirect: Yellow
    *   403 Forbidden: Red
    *   404 Not Found: Gray
*   **Wordlist handling:** Read the wordlist file and send HTTP requests for each entry in the wordlist [Reference 3]().
*   **Output:** Display the results in the terminal with appropriate colors, and save them to an output file if specified. To easily copy and paste the output, use the following command structure:

    ```bash
    ./Rfuzz -u <target_url> -w <wordlist_file> -t <number_of_threads> -o <output_file>
    ```

    *Replace `<target_url>`, `<wordlist_file>`, `<number_of_threads>`, and `<output_file>` with your desired values.*

    OK, here's the "Steps to Fix and Run the Go Program" section formatted attractively for a GitHub README, focusing on copy-pastability:

### Steps to Fix and Run the Go Program

If you encounter issues running the Go program, follow these steps to resolve common problems related to line endings and compilation.

1.  **Convert Line Endings**

    *   Sometimes, files created or edited on Windows systems have different line endings than those expected on Unix-like systems (like Linux or macOS). This can cause issues when compiling or running the program.
    *   Use the `dos2unix` command to convert the file's line endings from Windows-style (CRLF) to Unix-style (LF):

        ```bash
        dos2unix Rfuzz.go
        ```

        *If `dos2unix` is not installed, you can install it using:*

        ```bash
        sudo apt install dos2unix  # For Debian/Ubuntu
        sudo yum install dos2unix  # For CentOS/RHEL/Fedora
        ```
2.  **Compile and Run the Program**

    *   Once the line endings are fixed, you can compile and run your Go program.
    *   **Compile the program:**

        ```bash
        go build Rfuzz.go
        ```

        *This creates an executable file named `Rfuzz` (or `Rfuzz.exe` on Windows).*
    *   **Run the compiled executable:**

        ```bash
        ./Rfuzz
        ```

    *   Alternatively, you can run the program directly without explicitly compiling it first, using the `go run` command:

        ```bash
        go run Rfuzz.go
        ```

        *This command compiles and runs the program in one step.  It's often convenient for development, but for deployment, compiling the executable is generally preferred.*
 
  

    
 
