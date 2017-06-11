--with Model; use Model;
with Ada.Strings.Unbounded;
with Ada.Strings.Bounded;
with Ada.Text_IO;use Ada.Text_IO;

package body symulacja is
    procedure cls is
    begin
        Ada.Text_IO.Put(ASCII.ESC & "[2J");
    end;

    procedure start is
    begin
        Put_Line("Hi!");
    end start;
end symulacja;
