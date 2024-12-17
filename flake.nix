{
  description = "Advent of Code Go development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          go                # Latest Go compiler and tools
          gopls            # Go language server
          delve           # Go debugger
          go-tools        # Additional Go tools like staticcheck
        ];

        shellHook = ''
          echo "🎄 Go Development Environment for Advent of Code"
          echo "Go version: $(go version)"
          zsh
        '';
      };
    };
}
