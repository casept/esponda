let
  sources = import ./nix/sources.nix;
  nixpkgs = import sources.nixpkgs { };
  vgo2nix = import sources.vgo2nix { };
  niv = import sources.niv { };
in nixpkgs.mkShell {
  buildInputs = [
    # Go dev tools
    nixpkgs.go
    nixpkgs.go-outline
    nixpkgs.gocode-gomod
    nixpkgs.gopkgs
    nixpkgs.go-outline
    nixpkgs.go-symbols
    #nixpkgs.guru
    #nixpkgs.gorename
    nixpkgs.gotests
    nixpkgs.gomodifytags
    nixpkgs.impl
    #nixpkgs.fillstruct
    #nixpkgs.goplay
    nixpkgs.gopls
    nixpkgs.delve
    nixpkgs.godef
    #nixpkgs.goreturns
    nixpkgs.golint
    vgo2nix
    niv.niv

    nixpkgs.imagemagick
    nixpkgs.statik
  ];
}
