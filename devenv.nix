{
  pkgs,
  lib,
  config,
  inputs,
  ...
}: {
  env.CGO_ENABLED = 1;

  packages = with pkgs; [
    gcc
    glib
    pkg-config
    cairo
    gobject-introspection
    golangci-lint
    golangci-lint-langserver
    graphene
  ];

  # https://devenv.sh/languages/
  languages = {
    go = {
      enable = true;
      package = pkgs.go_1_22;
    };
  };

  # https://devenv.sh/pre-commit-hooks/
  # pre-commit.hooks.shellcheck.enable = true;

  # See full reference at https://devenv.sh/reference/options/
}
