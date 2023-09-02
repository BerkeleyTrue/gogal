{
  description = "a web gallary browser for your photos";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    nix-filter.url = "github:numtide/nix-filter";
  };

  outputs = { self, nixpkgs, flake-utils, nix-filter }:
    (flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = import nixpkgs {
            inherit system;
          };
        in
        {
          formatter = pkgs.alejandra;
          packages.default = pkgs.callPackage ./. { inherit pkgs nix-filter; };
          devShells.default = import ./shell.nix { inherit pkgs; };
        })
    );
}
