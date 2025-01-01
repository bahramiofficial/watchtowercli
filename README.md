It looks like you’re repeating the fixed version I provided earlier. Here’s the same markdown text with a bit of cleanup for better clarity and formatting:

# Watchtower CLI

Watchtower CLI is a command-line tool designed for syncing programs with a database. Built with Go, it allows users to interact with a program synchronization feature via a custom terminal alias. The tool can be used to automate syncing programs to the database with ease.

## Table of Contents
- [Installation](#installation)
- [Setup](#setup)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install Watchtower CLI, follow these steps:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/watchtowercli.git
   cd watchtowercli

	2.	Install dependencies:

go mod tidy


	3.	Build the project (optional):

go build



Setup
	1.	Open terminal and run the following to add an alias to your shell configuration:

nano ~/.zshrc


	2.	Add the following line at the end of the file:

alias watch_sync_programs="go run /Users/mrpit/Documents/GitHub/watchtowercli/cmd/main.go syncprograms"


	3.	Save and exit the file. Then, run:

source ~/.zshrc


	4.	Now, you can run the program using the alias:

watch_sync_programs



Usage

Once the setup is complete, you can run the Watchtower CLI tool from the terminal using the defined alias. This will trigger the syncing process of programs to the database.

Running Sync Program

To initiate syncing, simply run:

watch_sync_programs

This will automatically run the syncprograms function in the Watchtower CLI.

Examples

Example 1: Sync Programs

Run the following command to sync programs:

watch_sync_programs

Example 2: Customizing the Sync Path

To change the directory where the JSON files are located, you can provide the path like this:

watch_sync_programs --path /path/to/your/json/files

Contributing

We welcome contributions! To get involved:
	1.	Fork the repository on GitHub.
	2.	Clone your fork to your local machine:

git clone https://github.com/yourusername/watchtowercli.git


	3.	Create a branch for your feature or bugfix:

git checkout -b feature-branch


	4.	Commit your changes:

git commit -m "Add new feature or fix bug"


	5.	Push the branch to your fork:

git push origin feature-branch


	6.	Submit a pull request on GitHub.

Please make sure to write tests and include relevant documentation for your contributions.

License

Watchtower CLI is licensed under the MIT License. See the LICENSE file for more information.

### Adjustments:
- Corrected any lingering formatting errors.
- Used proper bullet points and spacing.
- Clarified some of the steps in the setup process. 

This version should now be clean and easy to follow.