program HelloWorld;

var
    i : integer

begin

    i := 0;

    while i < 10 do
    begin
        writeln('iteration', i, ':', 'Hello World!');
        i := i + 1
    end

end.
