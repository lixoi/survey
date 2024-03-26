{name ? "lixoi/fluent-bit-distroless", tag ? "v1" }:
let 
	system = "x86_64-linux";
	pkgs = import <nixpkgs> {inherit system; };
	pkgsCross.aarch64-multiplatform.pkgsStatic = import <nixpkgs> {inherit system; };
	staticPkgs = pkgsCross.aarch64-multiplatform.pkgsStatic; 
 in 
  staticPkgs.dockerTools.buildImage {
	inherit name tag;
	copyToRoot = pkgs.buildEnv {
		name = "fluent-bit";
		pathsToLink = ["/" "/bin"];
		paths = [staticPkgs.fluent-bit ./conf];
	};
	config = {
		Env = [
		];
   		Cmd = [
		];
		Entrypoint = ["fluent-bit"];
 	};
 }
