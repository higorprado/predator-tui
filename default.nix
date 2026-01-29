{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoModule rec {
  pname = "predator-tui";
  version = "0.1.0";

  src = ./.;

  vendorHash = lib.fakeHash;

  meta = with pkgs.lib; {
    description = "Acer Predator Control Center TUI";
    homepage = "https://github.com/higorprado/predator-tui";
    license = licenses.mit;
  };
}