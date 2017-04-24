with Ada.Text_IO; use Ada.Text_IO;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;


procedure z1 is 

    type Stacja is record
        Id_Stacji : Integer;
        Czas_Zwrotnicy : Integer;
        Czy_Wolna : Boolean;
    end record;

    subtype positive is integer range 1..integer'last;
    type Stacje is array(positive range<>) of Stacja;

    Buffer_Size : Constant := 2000;
    Last : Natural;
    Item : String(1 .. Buffer_Size);

    File : File_Type;
    Value : Unbounded_String := Null_Unbounded_String;

    Ilosc_Stacji : Integer;
    Array_Stacje : Stacje(1 .. 8);
begin
    Open(File => File,
         Mode => In_File,
         Name => "../data.txt");
    Get_Line(File,Item, Last);
    Append(Source => Value, New_Item => Item (1 .. 1));
    Ilosc_Stacji := Integer'Value(To_String(Value));
    
    for I in Array_Stacje'Range loop
        Put(Integer'Image(I));
    end loop;

    loop
        exit when End_Of_File(File);
        Value := Ada.Strings.Unbounded.To_Unbounded_String(Get_Line(File));
--        Put(Ada.Strings.Unbounded.To_String(Value));
--        New_Line;
    end loop;

    Close(File);
    --Put(Integer'Image(Stacja_1.Id_Stacji));
end z1;
