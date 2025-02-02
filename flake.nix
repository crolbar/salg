{
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = inputs: let
    systems = ["x86_64-linux"];
    eachSystem = inputs.nixpkgs.lib.genAttrs systems;
    pkgsFor = eachSystem (system: import inputs.nixpkgs {inherit system;});
  in {
    packages = eachSystem (
      system: let
        pkgs = pkgsFor.${system};
      in {
        default = pkgs.buildGoModule {
          pname = "salg";
          version = "0.1";
          src = ./.;

          vendorHash = "sha256-Koc5Z//NrrPyFVR+XchoL2DIQc9/vUCr9XUDPucIFOA=";
        };
      }
    );
  };
}
