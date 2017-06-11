  use HTTP::Daemon;
  use HTTP::Status;  
  #use IO::File;

  my $d = HTTP::Daemon->new(
           LocalAddr => 'localhost',
           LocalPort => 4321,
       )|| die;
  
  print "Please contact me at: <URL:", $d->url, ">\n";


  while (my $c = $d->accept) {
      while (my $r = $c->get_request) {
          if ($r->method eq 'GET') {
              
              $file_s= "./index.html";    # index.html - jakis istniejacy plik
              #$c->send_file_response($file_s);
              $c->send_basic_header;
              $c->send_header(
                'Content-Type' => 'image/jpeg',
                'Content-Length' => '56360',
                'Accept-Ranges'  => 'bytes',
              );
              $c->send_crlf;

          }
          else {
              $c->send_error(RC_FORBIDDEN)
          }

      }
      $c->close;
      undef($c);
  }
