<!DOCTYPE html>

<html lang="en">
{{ partial "head.html" . }}

<body>
  <div id="site-wrapper">
    {{ partial "header.html" . }}

    <div class="jumbotron">
      <div class="container">
        <img src="/images/nacelle.png" alt="Nacelle logo">
        <p>
          Go Microservice Framework
          <br />
          Get your Apps in Flight
        </p>
      </div>
    </div>

    <div class="container">
      <div class="row">
        <div class="col-md-4 col-sm-8 col-xs-12 feature-item">
          <i class="icon pull-left fa fa-server"></i>
          <p>
            <a href="/docs/core/process">Process Management</a>:
            Nacelle initializes and monitors internal processes, ensuring the application continues to report healthy.
          </p>
        </div>

        <div class="col-md-4 col-sm-8 col-xs-12 feature-item">
          <i class="icon pull-left fa fa-cogs"></i>
          <p>
            <a href="/docs/core/config">Configuration Management</a>:
            Nacelle populates user-defined configuration structus from the environment and provides basic type conversions and validation logic.
          </p>
        </div>

        <div class="col-md-4 col-sm-8 col-xs-12 feature-item">
          <i class="icon pull-left fas fa-project-diagram"></i>
          <p>
            <a href="/docs/core/service">Dependency Injection</a>:
            Nacelle injects dependencies into user-defined process structs from an application-level service container.
          </p>
        </div>
      </div>

      <div class="row">
        <div class="col-md-4 col-sm-8 col-xs-12 feature-item">
          <i class="icon pull-left fas fa-stream"></i>
          <p>
            <a href="/docs/core/log">Ubiquitous Logging</a>:
            Nacelle encourages opinionated structured logging to stderr at all layers of the application.
          </p>
        </div>

        <div class="col-md-4 col-sm-8 col-xs-12 feature-item">
          <i class="icon pull-left fas fa-file-code"></i>
          <p>
            <a href="/docs/libraries">Base Libraries</a>:
            Nacelle provides skeleton processes suitable for HTTP servers, gRPC servers, AWS Lambda servers, and generic worker processes.
          </p>
        </div>

        <div class="col-md-4 col-sm-8 col-xs-12 feature-item">
          <i class="icon pull-left far fa-file-code"></i>
          <p>
            <a href="/docs/frameworks">Frameworks</a>:
            Nacelle forms the base for the Chevron HTTP server framework and the Scarf gRPC server framework.
          </p>
        </div>
      </div>

      <div class="row top-hr">
        <div class="col-md-12"><hr></div>
      </div>

      <div class="row open-source">
        <div class="col-md-12">
          <h1><i class="fab fa-github"></i></h1>
          <p>
            Nacelle is 100% open source.
            <br />
            All components are available under the <a href="https://opensource.org/licenses/MIT">MIT License</a> on <a href="https://github.com/go-nacelle">GitHub</a>.
          </p>
        </div>
      </div>
    </div>
  </div>

  {{ partial "footer.html" . }}
</body>
</html>
