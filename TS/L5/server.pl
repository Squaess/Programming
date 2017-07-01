use HTTP::Daemon;
use HTTP::Status;

#domyślny folder z plikami
$SOURCE_PATH = '/home/aedd/www';

my $d = HTTP::Daemon->new(
		ReuseAddr => 1, 
		LocalAddr => 'localhost',
		LocalPort => 4321,
	) || die;

print "Please contact me at : <URL:", $d->url,">\n";

while (my $c = $d->accept) {
	while (my $r = $c->get_request){
		if ($r->method eq 'GET'){
			$file_s = $r->uri;

			if ($file_s eq "/") {
				$file_s = "/index.html";
			}

			$c->send_file_response($SOURCE_PATH.$file_s);
		}
		else {
			$c->send_error(RC_FORBIDDEN)
		}
	}
	$c->close;
	undef($c);
}