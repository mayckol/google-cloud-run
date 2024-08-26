# Cloud Run Application

This is a Go-based application designed to be deployed on Google Cloud Run. The application provides weather data based on the given zip code.

## Getting Started

### Prerequisites

Ensure you have Docker and Docker Compose installed on your local machine.

### Environment Variables

The application requires specific environment variables to function correctly. These variables can be provided either through a `.env` file or directly in the `docker-compose.yml` file. Below is an explanation of both methods.

#### Method 1: Using a `.env` File

The recommended approach is to create a `.env` file in the root directory of your project. This file should contain all the necessary environment variables. Docker Compose will automatically load these variables into the container.

Here is an example of what your `.env` file might look like:

```dotenv
VIA_CEP_API_URL=https://viacep.com.br/ws
WEATHER_API_URL=http://api.weatherapi.com/v1/current.json
WEATHER_API_KEY=your_key_here
PORT=8080
```

#### Method 2: Directly in `docker-compose.yml`
Alternatively, you can uncomment the environment section in the docker-compose.yml file and directly set the environment variables there. Here’s how you can do it:
```yml
app:
  build:
    context: .
    dockerfile: Dockerfile
  container_name: cloudrun_app
  ports:
    - "8080:8080"
  environment:
    - VIA_CEP_API_URL=https://viacep.com.br/ws
    - WEATHER_API_URL=http://api.weatherapi.com/v1/current.json
    - WEATHER_API_KEY=your_key_here
    - PORT=8080
  command: ["./cloudrun"]
```

### Important
- If you choose to use the .env file, make sure the env_file section in docker-compose.yml is uncommented, and the environment section is commented out.
- If you prefer to define the variables directly in docker-compose.yml, make sure the environment section is uncommented, and the env_file section is commented out.
- Don’t forget to replace your_key_here with your actual API key for the weather service.

### Running the Application

To build and start the application using Docker, run the following command:

```bash
docker-compose up -d
```
This will build the Docker image, start the container, and run the application in detached mode.

### Running Tests
To run the tests inside the Docker container, use the following command:

```shell
docker-compose -f docker-compose.yml run --rm cloud_run_test
```
This command will execute the test suite and then remove the test container.

### Accessing the Remote Application
The application is also deployed remotely and can be accessed at:

**Remote URL [https://google-cloud-run-hfjjq7jqqa-rj.a.run.app/](https://google-cloud-run-hfjjq7jqqa-rj.a.run.app/)**
### API Endpoints

200 Example  
To get a successful weather result, use the following endpoint:
```
GET https://google-cloud-run-hfjjq7jqqa-rj.a.run.app/weather/91030-910
```
404 Example  
This endpoint returns a 404 error when the requested resource is not found:
```
GET https://google-cloud-run-hfjjq7jqqa-rj.a.run.app/weather/19999-999
```
422 Example
This endpoint returns a 422 error when the input is invalid:
```
GET https://google-cloud-run-hfjjq7jqqa-rj.a.run.app/weather/199992332-999
```
### Notes
- Ensure you have the necessary environment variables set up in the .env file if required by the application.
- The application is designed to handle various error cases and return appropriate HTTP status codes based on the input.
