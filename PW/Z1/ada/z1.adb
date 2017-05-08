with Ada.Text_IO; use Ada.Text_IO;
with Ada.Integer_Text_IO;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;
with Conveyance3; use Conveyance3;


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

    Hummer  : TRANSPORT;
    Limo    : CAR;
    Chevy   : CAR;
    Dodge   : TRUCK;
    Ford    : TRUCK;

begin
    Set_Values(Hummer, 4, 5760.0);

    Set_Values(Limo, 8);
    Set_Values(TRANSPORT(Limo), 4, 3750.0);

    Set_Values(Chevy, 5);
   Set_Values(TRANSPORT(Chevy), 4, 2560.0);

   Set_Values(Dodge, 6, 4200.0, 3, 1800.0);
   Set_Values(Ford, 4, 2800.0, 3, 1100.0);

   Put("The Ford truck has");
   Ada.Integer_Text_IO.Put(Get_Wheels(Ford), 2);
   Put(" wheels, and can carry");
   Ada.Integer_Text_IO.Put(Get_Passenger_Count(Ford), 2);
   Put(" passengers.");
   New_Line;

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
