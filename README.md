# zstash
================

**Secure File Encryption and Storage**

zstash is a Rust-based project designed to provide a secure and efficient way to encrypt and store sensitive files. This project aims to offer a simple, yet robust solution for protecting sensitive data from unauthorized access.

## Key Features
---------------

* **Secure Encryption**: zstash utilizes industry-standard encryption algorithms to ensure the confidentiality and integrity of your files.
* **Efficient Storage**: The project is designed to optimize storage space while maintaining the security and accessibility of your encrypted files.
* **Easy-to-Use Interface**: zstash provides a user-friendly interface for encrypting, decrypting, and managing your files.

## Installation Instructions
-------------------------

To install zstash, follow these steps:

1. **Install Rust**: Make sure you have Rust installed on your system. You can download and install Rust from the official [Rust installation page](https://www.rust-lang.org/tools/install).
2. **Clone the Repository**: Clone the zstash repository using Git: `git clone https://github.com/your-username/zstash.git`.
3. **Build the Project**: Navigate to the project directory and run `cargo build` to build the zstash executable.

## Usage Examples
-------------

### Encrypting a File

To encrypt a file using zstash, run the following command:
```bash
./zstash encrypt -i input_file.txt -o encrypted_file.zst
```
Replace `input_file.txt` with the path to the file you want to encrypt, and `encrypted_file.zst` with the desired output file path.

### Decrypting a File

To decrypt a file using zstash, run the following command:
```bash
./zstash decrypt -i encrypted_file.zst -o decrypted_file.txt
```
Replace `encrypted_file.zst` with the path to the encrypted file, and `decrypted_file.txt` with the desired output file path.

## Project Structure Overview
---------------------------

The zstash project consists of the following directories and files:

* `src`: Contains the Rust source code for the zstash executable.
* `tests`: Includes unit tests and integration tests for the zstash project.
* `Cargo.toml`: The project's Cargo configuration file.
* `.idea`: IntelliJ IDEA project configuration files.

## Dependencies and Requirements
-----------------------------

* Rust 1.58 or later
* Cargo 1.58 or later
* OpenSSL 1.1.1 or later (for encryption)

## Contributing Guidelines
------------------------

Contributions to zstash are welcome! If you'd like to contribute, please follow these guidelines:

1. Fork the zstash repository on GitHub.
2. Create a new branch for your feature or bug fix.
3. Commit your changes with a clear and concise commit message.
4. Open a pull request to merge your changes into the main branch.

## License Information
----------------------

zstash is licensed under the [MIT License](https://opensource.org/licenses/MIT). You are free to use, modify, and distribute zstash under the terms of this license.

Note: This README is a template and should be updated to reflect the actual project details and features.

# zStash
================

`zStash` is a command-line tool for securely encrypting and decrypting files using AES-GCM encryption with password-based key derivation (PBKDF2). This tool provides a simple and secure way to protect sensitive files from unauthorized access.

## Key Features
-------------

*   **File Encryption and Decryption**: `zStash` uses AES-GCM encryption to secure files, ensuring confidentiality and integrity of data.
*   **Password-Based Key Derivation**: The tool uses PBKDF2 to derive a strong encryption key from a user-provided password, making it resistant to brute-force attacks.
*   **Secure Storage**: `zStash` stores encrypted files in a designated directory, ensuring that sensitive data is kept separate from other files.

## Installation Instructions
-------------------------

To install `zStash`, follow these steps:

### Using Cargo (Rust Package Manager)

1.  Install Rust and Cargo from the official Rust installation page: <https://www.rust-lang.org/tools/install>
2.  Clone the `zStash` repository using the following command:

    ```bash
git clone https://github.com/Satuiro/zstash.git
```
3.  Navigate to the cloned repository and run the following command to build and install `zStash`:

    ```bash
cargo build --release
cargo install --path .
```

### Manual Installation

1.  Clone the `zStash` repository using the following command:

    ```bash
git clone https://github.com/Satuiro/zstash.git
```
2.  Navigate to the cloned repository and run the following command to build `zStash`:

    ```bash
cargo build --release
```
3.  Copy the compiled `zStash` executable to a directory in your system's PATH, such as `/usr/local/bin`.

## Usage Examples
--------------

Here are some basic usage examples for `zStash`:

### Encrypting a File

```bash
zStash encrypt --input example.txt --output encrypted_example.txt
```

### Decrypting a File

```bash
zStash decrypt --input encrypted_example.txt --output decrypted_example.txt
```

### Listing Available Commands

```bash
zStash --help
```

## Project Structure Overview
---------------------------

The `zStash` project is organized into the following directories and files:

*   `Cargo.toml`: The Rust package file, containing metadata and dependencies.
*   `Cargo.lock`: The generated lock file, ensuring consistent dependencies.
*   `src/lib.rs`: The main library file, containing the encryption and decryption logic.
*   `src/main.rs`: The main executable file, handling command-line arguments and executing the encryption and decryption commands.

## Dependencies and Requirements
------------------------------

`zStash` relies on the following dependencies:

*   `clap` for command-line argument parsing
*   `aes-gcm` for AES-GCM encryption
*   `pbkdf2` for password-based key derivation

These dependencies are automatically managed by Cargo, the Rust package manager.

## Contributing Guidelines
----------------------

Contributions to `zStash` are welcome and encouraged. If you'd like to contribute, please follow these steps:

1.  Fork the `zStash` repository on GitHub.
2.  Create a new branch for your feature or bug fix.
3.  Commit your changes with clear and descriptive commit messages.
4.  Open a pull request to the main `zStash` repository.

## License Information
---------------------

`zStash` is licensed under the MIT License or the Apache License 2.0. You may choose to use either license at your discretion. For more information, please refer to the `LICENSE` file in the repository.