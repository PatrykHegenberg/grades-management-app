# Grades Management Application

This application is designed to manage student grades, allowing users to input grades, calculate averages, and export results as a PDF. It is built using Go and Wails for the backend and Svelte for the frontend.

## Features

- **Add Grades**: Users can add grades for "Hörverstehen" and "Leseverstehen". If only one partial score is required, the second can be left blank and the percentage of the first set to 100
- **Grade Overview**: Displays an overview of grades (1-6) with the count of each grade.
- **PDF Export**: Allows users to export the grades to a PDF file.
- **Interactive User Interface**: Built with Svelte, providing a responsive and user-friendly experience.

## Installation

### Prerequisites

Make sure you have the following software installed:

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Node.js](https://nodejs.org/) (for building the frontend)
- [Wails](https://wails.io/docs/gettingstarted/installation) (install Wails CLI using `go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

### Clone the Repository

```bash
git clone https://github.com/PatrykHegenberg/grades-management-app.git
cd grades-management-app
```

### Start the Application in dev mode

From main directory run the Wails dev command:

```bash
wails dev
```

## Usage

- Set Maximum Points: Enter maximum points and weights for grades.
- Add Grade: Add a grade by entering first name, last name, hv grade, and lv grade.
- View Grade Overview: The number of grades for each score will be automatically updated.
- Export: Choose a location and export the grades as a PDF.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contribution

Contributions are welcome! Please open an issue or pull request.
