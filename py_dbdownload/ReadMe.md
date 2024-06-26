# MongoDB Backup and Restore Scripts

This repository contains Python scripts for backing up and restoring MongoDB collections to and from JSON files. The scripts are designed to be modular and easy to use, ensuring that the logic is not repeated and functions are used to streamline the process.

## Features

- **Backup MongoDB Collections**: Export MongoDB collections to JSON files, saving them in a specified directory.
- **Restore MongoDB Collections**: Import data from JSON files into MongoDB collections.
- **Modular Design**: The scripts are organized into reusable functions, making them easy to maintain and extend.

## Requirements

- Python 3.x
- `pymongo` library

## Setup

1. **Install Python 3.x**: Ensure that Python 3.x is installed on your system. You can download it from the [official Python website](https://www.python.org/downloads/).

2. **Install pymongo**: Install the `pymongo` library using pip:
    ```bash
    pip install pymongo
    ```

## Configuration

1. **MongoDB URI**: Set your MongoDB URI in the script. The default URI is `mongodb://localhost:27017`.

2. **Database Names**: Specify the names of the databases you want to back up and restore.

3. **Directories**: Define the directories for storing backup files and loading restore files. The default directories are:
    - Backup directory: `~/Documents/mongopy/mongodb_backup`
    - Restore directory: `~/Downloads/mongodb_backup`

## Usage

### Backup MongoDB Collections

To backup MongoDB collections to JSON files:

1. Ensure your MongoDB server is running.
2. Run the script to back up the collections:
    ```bash
    python main.py
    ```
3. The JSON files will be saved in the specified backup directory.

### Restore MongoDB Collections

To restore MongoDB collections from JSON files:

1. Ensure your MongoDB server is running.
2. Place the JSON files in the specified restore directory.
3. Run the script to restore the collections:
    ```bash
    python main.py
    ```
4. The data will be inserted into the specified MongoDB collections.

## Notes

- The backup script will convert `ObjectId` to strings for JSON serialization.
- The restore script can handle both single and multiple documents in the JSON files.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## Contact

For any inquiries, please contact [your_email@example.com](mailto:your_email@example.com).
