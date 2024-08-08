# My Go Application

## Overview

This Go application provides functionality for handling image analysis, product information retrieval, product search, and recipe generation. It utilizes OpenAI's image analysis capabilities and MongoDB for data storage.

## Setup

1. **Clone the repository:**

    ```sh
    git clone https://github.com/tarakanillius/backendYuka.git
    cd your-repository
    ```

2. **Set up environment variables:**

    Create a `.env` file in the root directory with the following content:

    ```env
    MONGODB_URI=mongodb://mongo:27017
    MONGODB_NAME=your_database_name
    OPENAI_API_KEY=your_openai_api_key
    ```

3. **Build and run the application:**

    ```sh
    docker-compose up --build
    ```

4. **Access the application:**

    The application will be available at `http://localhost:8080`.

## API Endpoints

- **POST /image/analyze**: Analyze an image and extract keywords.
- **GET /product/{id}**: Retrieve a product by its ID.
- **GET /search**: Search for products based on keywords.
- **POST /generate-receipt**: Generate a recipe based on a list of products.
- **POST /generate-recommendations**: Generate recommendations based on a list of products.

## License

This project is licensed under the MIT License.
