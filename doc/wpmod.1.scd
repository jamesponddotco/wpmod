wpmod(1)

# NAME

wpmod - WordPress file permissions hardening tool

# SYNOPSIS

*wpmod* [OPTIONS]

# DESCRIPTION

*wpmod* is a command-line tool designed to set WordPress file and directory
permissions according to the recommendations by the WordPress core team. It also
provides options to set strict permissions for sensitive files.

# OPTIONS

*-p*, *--path*
	Specify the path to the WordPress installation. This option is required.

*-u*, *--user*
	Specify the user for file ownership. This option is required.

*-g*, *--group*
	Specify the group for file ownership. This option is required.

*-s*, *--strict*
	Whether to set strict permissions for sensitive files. The "wp-config.php"
	file and the "/mu-plugins/" directory are considered sensitive.

*-h*, *--help*
	Show help message and quit.

*-v*, *--version*
	Show version number and quit.

## EXAMPLES

*Example 1. Set WordPress file permissions according to the recommendations.*
	$ wpmod -p '/srv/example.com/' -u 'www-data' -g 'www-data'

*Example 2. Set strict permissions for sensitive files.*
	$ wpmod -p '/srv/example.com/' -u 'www-data' -g 'www-data' -s

# SEE ALSO

*chmod*(1), *chown*(1)

# AUTHORS

Maintained by James Pond <james@cipher.host>.
