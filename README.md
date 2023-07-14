![Photo by Brett Ritchie on Unsplash](https://user-images.githubusercontent.com/2342458/202707249-359f3a1b-d61c-4d3b-a8b7-7a284f8f9842.png)

# Kinsta - Hello World - Go
An example of how to set your Go application up to enable deployment on Kinsta App Hosting services.

---
Kinsta is a developer-centric cloud host / PaaS. We’re striving to make it easier for you to share your web projects with your users. Focus on coding and building, and we’ll take care of deployment and provide fast, scalable hosting. + 24/7 expert-only support.

- [Start your free trial](https://kinsta.com/signup/?product_type=app-db)
- [Application Hosting](https://kinsta.com/application-hosting)
- [Database Hosting](https://kinsta.com/database-hosting)

## Dependency Management

During the deployment process Kinsta will automatically install dependencies defined in your go.mod file.

## Web Server Setup

### Port
Kinsta automatically sets the `PORT` environment variable. You should **not** define it yourself and you should **not** hard-code it into the application. Use
`os.Getenv("PORT")` in your code when referring to the server port.

### Start Command
When deploying an application, Kinsta will automatically create a web process.

## Watch How to Set Up a Go Application on Kinsta
[![Watch the video](https://img.youtube.com/vi/AftCe2tzILU/maxresdefault.jpg)](https://www.youtube.com/watch?v=AftCe2tzILU)
