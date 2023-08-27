{ pkgs }:
let
  # update the vendorSha256 of the default package
  update-vendor-sha = pkgs.writeShellScriptBin "update-vendor-sha" ''

  set -exuo pipefail

  failedbuild=$(nix build --impure 2>&1 || true)
  # echo "$failedbuild"
  checksum=$(echo "$failedbuild" | awk '/got:.*sha256/ { print $2 }')
  echo -n "\n\nchecksum: $checksum"
  # do nothing if no checksum was found
  if [ -z "$checksum" ]; then
    exit 0
  fi
  sed -i -e "s|vendorSha256 = \".*\"|vendorSha256 = \"$checksum\"|" ./default.nix
'';
in
pkgs.mkShell {
  name = "gogal";
  packages = with pkgs; [
    go
    gopls # language server
    gotools
    go-tools
    air # live reload

    update-vendor-sha
    nodePackages.vscode-langservers-extracted # html/css language server
  ];

  # enter zsh on startup
  shellHook = ''
    zsh
    exit
  '';
}
