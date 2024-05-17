# Goose usage
- in order to keep it tidy I like using Goose to migrate my database with the cli tool [Goose](https://github.com/pressly/goose)

# Stuff i struggled to figure our 
- After installing the Goose cli tool you can use it in many diffrent ways
- My approch is to set the env file like in the example.env and source the env .
Just after the -env source you can use goose with the env files before that you will always 
have to put the dbstring dbtype and so on
