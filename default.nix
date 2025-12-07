{ pkgs ? import
    (fetchTarball {
      name = "jpetrucciani-2025-12-07";
      url = "https://github.com/jpetrucciani/nix/archive/daabd563787e50d3cb1eab13abde55f84ac0b057.tar.gz";
      sha256 = "1r9db59cdys74jq9c5q3dg8q2wwd7w2pdbdn9r89zmyyfsiymaql";
    })
    { }
}:
let
  name = "forge";

  tools = with pkgs; {
    cli = [
      jfmt
      nixup
    ];
    go = [
      go
      go-tools
      gopls
    ];
    scripts = pkgs.lib.attrsets.attrValues scripts;
  };

  scripts = with pkgs; { };
  paths = pkgs.lib.flatten [ (builtins.attrValues tools) ];
  env = pkgs.buildEnv {
    inherit name paths; buildInputs = paths;
  };
in
(env.overrideAttrs (_: {
  inherit name;
  NIXUP = "0.0.10";
})) // { inherit scripts; }
