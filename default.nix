{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoModule rec {
  pname = "predator-tui";
  version = "0.1.0";

  src = ./.;

  vendorHash = "sha256-lKSr05aeK+HBxJKIbBPSesYpokf6D2Yol8p4OHHjNQ8=";

  meta = with pkgs.lib; {
    description = "Acer Predator Control Center TUI";
    homepage = "https://github.com/higorprado/predator-tui";
    license = licenses.mit;
  };
}