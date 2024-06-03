# File Upload App

## Description
This is a simple file upload application built with Go, Gin, and Bootstrap. It allows users to upload files and view them in either thumbnail or list mode. Additionally, a QR code with the local IP address is displayed on the home page.

## Features
- Upload multiple files
- View uploaded files as thumbnails or in a list with file size
- Display QR code with local IP address
- Mobile-friendly UI with Bootstrap

## Setup
1. Clone the repository:
    ```bash
    git clone <repository_url>
    ```
2. Change directory to the project folder:
    ```bash
    cd file-upload-app
    ```
3. Install dependencies:
    ```bash
    go mod tidy
    ```
4. Create a configuration file (`config.yaml`) in the project root with the following content:
    ```yaml
    upload_path: "./uploads"
    ```
5. Run the application:
    ```bash
    go run main.go
    ```

## Usage
- Open a web browser and go to `http://localhost:8080`
- Use the upload page to upload files
- View the uploaded files on the home page

## Project Structure
- `main.go`: The main entry point of the application
- `config/`: Configuration-related files
- `handlers/`: Route handlers for the application
- `utils/`: Utility functions
- `templates/`: HTML templates

## Dependencies
- [Gin](https://github.com/gin-gonic/gin): Web framework for Go
- [Viper](https://github.com/spf13/viper): Configuration management for Go
- [go-qrcode](https://github.com/skip2/go-qrcode): QR code generation in Go
- [Bootstrap](https://getbootstrap.com/): Front-end framework for responsive web design

## License
This project is licensed under the MIT License.