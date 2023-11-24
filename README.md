# `wpmod`

`wpmod` is a simple tool written in Go that sets WordPress' file and
directory permissions to [the ones recommended by the WordPress core
team](https://wordpress.org/support/article/hardening-wordpress/#file-permissions).
The tool also improves permissions for the _wp-config.php_ file and the
_/mu-plugins/_ directory if `--strict` is used; testing is recommended.

I got tired of seeing plugins and users changing permissions and
breaking everything, so I wrote this as a solution. Set it to run every
_N_ hours with a cron job and you should be good to go.

## Installation

### From source

First install the dependencies:

- Go 1.21 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```bash
make
sudo make install
```

## Usage

```bash
$ wpmod --help
NAME:
   wpmod - harden WordPress' file permissions

USAGE:
   wpmod [global options] [arguments...]

VERSION:
   0.1.0

GLOBAL OPTIONS:
   --path value, -p value   path to the WordPress installation
   --user value, -u value   user for file ownership
   --group value, -g value  group for file ownership
   --strict, -s             enable strict file permission mode (default: false)
   --help, -h               show help
   --version, -v            print the version
```

See _wpmod(1)_ after installing for more information.

## Contributing

Anyone can help make `wpmod` better. Send patches on the [mailing
list](https://lists.sr.ht/~jamesponddotco/wpmod-devel) and report bugs
on the [issue tracker](https://todo.sr.ht/~jamesponddotco/wpmod).

You must sign-off your work using `git commit --signoff`. Follow the
[Linux kernel developer's certificate of
origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin)
for more details.

All contributions are made under [the GPL-2.0 license](LICENSE.md).

## Resources

The following resources are available:

- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/wpmod-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/wpmod-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/wpmod).

---

Released under the [GPL-2.0 license](LICENSE.md).
