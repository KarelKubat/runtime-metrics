#!/usr/bin/env perl

# Pre-push hook for Go projects that end up on github.
# 
# Each directory must have a doc.go.
# For each .go file, a _test.go must exist.
# In each directory, `go test` is run.
# In each directory, `godocdown` is run.

use strict;
use Cwd;
use File::Compare;
use File::Find;

my %dirs;
my %errs;

sub checkerrs {
    my $errfound = 0;
    for my $k (sort(keys(%errs))) {
	++$errfound;
	print STDERR ("pre-push error: $k\n");
    }
    if ($errfound) {
	print STDERR ("$errfound missing so far, stopping\n");
	exit(1);
    }
}

sub wanted {
    my $basefile = $_;
    my $dir = $File::Find::dir;
    my $file = $File::Find::name;

    # Keep out of dirs where you should not go.
    return if ($dir =~ m{.git});

    # Ignore _test.go files
    return if ($file =~ m{_test.go$});

    $dirs{$dir} = 1;

    # Directories must have a doc.go.
    $errs{"Missing $dir/doc.go"} = 1 unless (-f "doc.go");
	  
    # Go files must have a _test.go
    if ($file =~ m{.go$}) {
	my $testfile = $basefile;
	$testfile =~ s{.go$}{_test.go};
	my $fulltest = $file;
	$fulltest =~ s{.go$}{_test.go};
	$errs{"Missing $fulltest (test for $file)"} = 1 unless (-f $testfile);
    }
}

# Check all subdirs for _test.go and doc.go presence
find(\&wanted, ('.'));
checkerrs();

# In all directories, run `go test` and run `godocdown`.
my $readme_modified = 0;
my $wd = getcwd();
for my $d (sort(keys(%dirs))) {
    chdir($d) || die("Cannot cd into $d: $!\n");

    print("Running tests in $d\n");
    system("go test") and $errs{"Failed 'go test' in $d"} = 1;
    
    if (! -f "README.md") {
	$readme_modified++;
	print("Generating $d/README.md\n");	
	system("godocdown . > README.md")
	  and $errs{"failed 'godocdown' in $d"} = 1;
    } else {
	unlink("README.prev");
	rename("README.md", "README.prev")
	  or die("Failed to rename README.md to .prev: $!\n");
	if (system("godocdown . > README.md")) {
	    $errs{"failed 'godocdown' in $d"} = 1;
	} else {
	    if (compare("README.md", "README.prev")) {
		$readme_modified++;
		print("Regenerated $d/README.md\n");
	    }
	}
	unlink("README.prev") or die("Cannot unlink README.prev: $!\n");
    }	
    chdir($wd) || die("Cannot cd back into $wd: $!\n");
}
checkerrs();

if ($readme_modified > 0) {
    print STDERR ("$readme_modified README.md's were (re)generated.\n",
		  "Please re-run 'git commit' and 'git push'.\n",
		  "Sorry for the inconvenience.\n");
    exit(1);
}
