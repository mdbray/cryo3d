##Development:
*Questions? Get in contact with:*
- Michelle Bray 		michelle.bray@colorado.edu
- Ali Hakimi			ali.hakimi@colorado.edu
- Annie Kelly			annie.kelly@colorado.edu
- Chris Gray			chris.gray@colorado.edu
- Lincoln Samelson		lincoln.samelson@colorado.edu
- Olivia Abrant			olivia.abrant@colorado.edu

###Initial Setup
####Github
Install Git if you do not already have it.  Click [here](https://help.github.com/articles/set-up-git/#platform-all) for setup instructions for all operating systems.

Fork the Cryo3D repo on Github:
- Find the Cryo3D repository [here] (https://github.com/cryo3d/cryo3d)
- In the top right corner of the page, click '**Fork**'
- Having trouble? Visit [Github's site](https://help.github.com/articles/fork-a-repo/)

####Local
Create a `/workspaces` directory and cd into it

Clone your fork of the Cryo3D repo into your workspaces directory
```
$ git clone {URL_of_your_fork}
```

Use `$ git remote -v` to check for a list of remotes synced to this local repo.  Ensure origin links back to your own fork.

Set an upstream path to sync back with the master branch of the original Cryo3D repo (you will use your own fork with branching for development and create pull requests back into the "master" repo.  This will ensure safe, consistent code via code reviews)

To set the appropriate upstream, type the following into your terminal:
```
$ git remote add upstream https://github.com/cryo3d/cryo3d
```

Throughout development cycle, pull changes from upstream to stay up-to-date with "master."  Use the following commands:
```
$ git fetch upstream
$ git pull upstream
```
**Make sure to checkout a new branch (on your own fork) with each major story/feature**

####Docker
Install Docker.  Reference Docker's setup instructions [here](http://docs.docker.com/windows/started/).  Links to help with all OS's can be found from this page.

Keep in mind that virtualization must be enabled.  This means a little extra setup work for OSX users.

For a simple 'Hello World' program with Docker, navigate [here] (https://github.com/docker-library/hello-world)

A detail explanation of how this works can be found [here] (https://github.com/docker-library/docs/tree/master/hello-world)

For a list of all running containers:
```
$ docker ps -l
```

`$ echo $DOCKER_HOST` within the container to bind the docker host to exposed port (8080 in this case) 

To stop/remove all Docker containers:
```
$ docker stop $(docker ps -a -q)
$ docker rm $(docker ps -a -q)
```

####Working with Go
For the purposes of setting up a simple server, install Go.

A walkthrough for multiple OS's can be found [here] (https://golang.org/doc/install)

After local installation, set your $GOROOT, $PATH, and $GOPATH environment variables.

If you followed the steps above, make a folder inside workspaces that will store all your go code.  Set your GOPATH to point to this location

For example, `~/workspaces/gocode` (in my case)

From the root of your go directory (`/gocode` in my case), create `src`, `bin`, and `pkg` folders 

For the sake of avoiding naming collisions with Go package management, set up a directory structure like so:
```
$GOPATH/src/github.com/{github_username}
```

Create the appropriate folders.

Fork the hello_Docker branch from the master repo into your own.  Clone your copy of this branch into `$GOPATH/src/github.com/{github_username}`


A helpful walkthrough for doing this on OS X can be found [here] (http://www.raulmurciano.com/articles/setting-up-go-for-development/)

`cd` into root (base of the working directory).

To build an image using the Dockerfile (Note: the `-t` flag tags the resulting image with the name hello_Docker), run:
```
$ docker build -t hello_Docker .
```

Next, run a container from the resulting image with
```
docker run --publish 8080:8080 --name test --rm hello_Docker
```

The server is now up and running at `http://{$DOCKER_HOST}:8080/`.  Make sure to replace the last colon and following four digits in your `$DOCKER_HOST` with our exposed port number (`:8080`).

For now, "Hello world!" should be printed to the screen.




