{inputs, ...}: {
  perSystem = {pkgs, ...}: {
    packages.default = pkgs.buildGoModule {
      pname = "gogal";
      version = "0.1";
      pwd = ./.;
      src = let
        show-trace = true;
        source-files =
          inputs.nix-filter.lib.filter
          {
            root = ./.;
          };
      in (
        if show-trace
        then pkgs.lib.sources.trace source-files
        else source-files
      );

      subPackages = ["cmd/cli"];
      vendorSha256 = "sha256-4A5j3N+H0TMYWFiVZUr74m3kK5kAcfZuXTDM/j+H024=";
    };
  };
}
