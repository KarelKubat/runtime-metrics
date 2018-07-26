#!/usr/bin/env perl

use strict;

my $missing = 0;
my $sources = 0;
foreach my $f (glob('*.go')) {
  next if ($f =~ /_test.go$/);
  my $testf = $f;
  ++$sources;
  $testf =~ s/.go$/_test.go/;
  if (! -f $testf) {
    warn("Missing: $testf (for $f)\n");
    ++$missing;
  }
}

die("One or more test files not found\n") if ($missing);
die("No *.go sources here (are you in the right cwd?)\n") unless ($sources);

