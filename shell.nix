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

  watch-compile = pkgs.writeShellScriptBin "watch-compile" ''
    ${pkgs.concurrently}/bin/concurrently -n css,go \
      "${pkgs.nodejs_18}/bin/npx tailwindcss -i ./web/input.css -o ./web/public/css/output.css --watch"\
      "${pkgs.air}/bin/air"
  '';

  watch-tests = pkgs.writeShellScriptBin "watch-tests" ''
    ${pkgs.ginkgo}/bin/ginkgo watch -r -p
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
    ginkgo # testing framework

    # editor
    nodePackages.vscode-langservers-extracted # html/css language server
    nodePackages.typescript-language-server # typescript language server

    # css stuff
    nodejs_18

    # scripts
    update-vendor-sha
    watch-compile
    watch-tests
  ];

  # enter zsh on startup
  shellHook = ''
    zsh
    exit
  '';
}
