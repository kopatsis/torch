Torch

This is a very basic template for any Go Gin project. 

The optional frontend stack is the default HTML templating, PostCSS with Tailwind, and Minify. The command "npm build:all" will compile all minified JS and CSS files, including the minimum necessary Tailwind CSS based on what was used in all of the HTML files. No hot reloading with Air (yet)

The backend by default uses Gorm for MySQL for a database in addition to Redis for caching/other actions. The repository pattern is used. In this case, the actual data models are defined under the models package. Each entity's repository exists under repositories, where direct database interaction takes place. Under services is where the "business layer" activities take place, which use their respective repositories. Finally, service.go contains the main service repository, MainService. Here you can access each entity's service functions via mainService.User.ServiceFunc1(args). Additionally if needed, repository-level functions can be passed up to the service layer directly. 

No real routing has been created yet nor have specialty packages, such as Firebase, Sendgrid, Stripe, etc as those are to be added for the specific implementation. 

To use this project, use the single line below

```git clone git@github.com:kopatsis/torch.git new-folder-name --depth 1 && rm -rf new-folder-name/.git```

Just use your own name. The module will still be named Torch, but it can be changed.