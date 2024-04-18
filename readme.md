
# ascii-art-web

This project is a simple ASCII Art Generator implemented in Go. It exposes two endpoints that can be accessed via HTTP. The ASCII Art Generator takes an input string and converts it into a corresponding ASCII art.

## Authors

- Hasan Dhaif (hadhaif)
- Ahmed Alhamed (ahalhamed)

## Usage

### Prerequisites

Go (version 1.16 or later)

### How to Run

1. Clone the repository

   ``` bash
   git clone https://learn.reboot01.com/git/hadhaif/ascii-art-web.git
   ```

2. Navigate to the project directory

   ``` bash
   cd ascii-art-web
   ```

3. Run the server

      ``` bash
      go run .
      ```

   The server runs on port 8080 and has two endpoints:

- ```/``` : This is the main endpoint which is handled by ```MainHandler```.
- ```/ascii-art``` : This endpoint is handled by ```AsciiHandler``` and generates ASCII art.

You can access these endpoints in your web browser or using a tool like curl.

## Implementation Details

The ASCII Art Generator uses a simple algorithm to convert each character of the input string into a corresponding ASCII art. The algorithm works by mapping each character to a predefined ASCII art pattern. The patterns are stored in a file for quick access, there is a separate file for each font. The algorithm iterates over each character of the input string, retrieves the corresponding ASCII art pattern from the file of the selected font, and appends it to the output string. The output string is then returned as the final ASCII art.

- The ```StartLineCalc``` function is used to calculate the starting line in the file of a specific ASCII art pattern for a given character.
- The ```ReadLines``` function is used to read a specific number of lines from the file of a specific ASCII art pattern starting from a given starting line.
- The ```AppendAscii``` function is used to append two ASCII art patterns.
- The ```AsciiLine``` function is used to generate the ASCII art for a single line of the input string using all three previously mentioned functions.
