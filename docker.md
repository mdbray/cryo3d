##Development:
*For questions, reach out to Michelle*

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




