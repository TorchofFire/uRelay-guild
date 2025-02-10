![Status](https://img.shields.io/badge/status-alpha-red) ![Alpha](https://img.shields.io/badge/version-v0.0.0-orange) ![MIT License](https://img.shields.io/badge/License-MIT-yellow.svg) ![MySQL](https://img.shields.io/badge/database-MySQL-blue) 
# uRelay - Guild
> **uRelay** is a project dedicated towards providing an open-source solution for online community communication.

**uRelay - Guild** is the server software for the decentralized chat platform. It lets you run your own private or public hub where people can connect and communicate free from restriction.

## Features
 - 

## Getting Started With Development

### Prerequisites
 - Make sure you have [Go](https://go.dev/) installed on your machine.
 - Have a mysql server set up and manually run the migration found in `internal/database/migrations`. https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/
 > Manually running migrations is temporary. If this seems like too much of a hassle right now, please wait till beta.

### Installation
1. Clone the repository:

    ```bash
    git clone https://github.com/TorchofFire/uRelay-guild
    ```

2. Navigate to the project folder:

    ```bash
    cd uRelay-guild
    ```
3. Setup environment variables:
	Using `.env.example`, make a copy and rename it to `.env`.
	Modify the file to include your database credentials.

4. Run it:

    ```bash
    go run .
    ```
> Every time you make a change you must restart the server for the changes to take effect.

## Contributing

Feel free to contribute to this project by opening issues or pull requests.

## License 
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
