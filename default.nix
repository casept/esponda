{ buildGoModule, imagemagick, ... }:
buildGoModule {
  pname = "esponda";
  version = "master";
  src = ./.;
  vendorSha256 = null;
  modSha256 = "0gbgcmnqka4781mg02r5vw51yz1jc2h30psbsvcinxgvayj4cqpk";
  buildInputs = [ ];
  nativeBuildInputs = [ imagemagick ];
}
