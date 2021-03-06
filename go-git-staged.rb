# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class GoGitStaged < Formula
  desc "Run commands on files staged in git.\nFilter files using globs and pass them to their respective commands as arguments."
  homepage "https://github.com/iiroj/go-git-staged"
  version "0.1.1"
  license "MIT"
  bottle :unneeded

  if OS.mac? && Hardware::CPU.intel?
    url "https://github.com/iiroj/go-git-staged/releases/download/v0.1.1/go-git-staged_0.1.1_darwin_amd64.tar.gz"
    sha256 "cd59ee8a25e06f3db6029ffe070343f564da26c856a272302a3e97bc459868f1"
  end
  if OS.mac? && Hardware::CPU.arm?
    url "https://github.com/iiroj/go-git-staged/releases/download/v0.1.1/go-git-staged_0.1.1_darwin_arm64.tar.gz"
    sha256 "292ecbbafaf507be1ebd3de3366d903fa15ecfd95bc7fe02e2821b4347a5f1d5"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/iiroj/go-git-staged/releases/download/v0.1.1/go-git-staged_0.1.1_linux_amd64.tar.gz"
    sha256 "40637f2da54c99d2d9adfe04a02de96f67d7b9e706bb659f2a163246282c26a2"
  end
  if OS.linux? && Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
    url "https://github.com/iiroj/go-git-staged/releases/download/v0.1.1/go-git-staged_0.1.1_linux_arm64.tar.gz"
    sha256 "7541df8e0976c6bd8d446586af9534e5db82e304064ebba03821d2d6bdae3204"
  end

  depends_on "git"

  def install
    bin.install "go-git-staged"
  end
end
