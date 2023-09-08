{
  pkgs,
  nix-filter,
}:
pkgs.buildGoModule {
  pname = "gogal";
  version = "0.1";
  pwd = ./.;
  src = let
    show-trace = true;
    source-files =
      nix-filter.lib.filter
      {
        root = ./.;
      };
  in (
    if show-trace
    then pkgs.lib.sources.trace source-files
    else source-files
  );

  subPackages = ["cmd/cli"];
  vendorSha256 = "sha256-85oFD9RvutaEPOLtaGGF6vFxwDKQ/A3wN2JBLgESHNg=";
}
