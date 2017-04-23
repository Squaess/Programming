package Conveyance1 is 

   -- This is a very simple transportation type.
   type TRANSPORT is
      record
         Wheels : INTEGER;
         Weight : FLOAT;
      end record;

   procedure Set_Values(Vehicle_In : in out TRANSPORT; 
                        Wheels_In  : INTEGER;
                        Weight_In  : FLOAT);
   function Get_Wheels(Vehicle_In : TRANSPORT) return INTEGER;
   function Get_Weight(Vehicle_In : TRANSPORT) return FLOAT;


   -- This CAR type extends the functionality of the TRANSPORT type.
   type CAR is new TRANSPORT;

   function Tire_Loading(Vehicle_In : CAR) return FLOAT;

end Conveyance1;




package body Conveyance1 is

-- Subprograms for the TRANSPORT record type.
procedure Set_Values(Vehicle_In : in out TRANSPORT; 
                     Wheels_In  : INTEGER; 
                     Weight_In  : FLOAT) is
begin
   Vehicle_In.Wheels := Wheels_In;
   Vehicle_In.Weight := Weight_In;
end Set_Values;

function Get_Wheels(Vehicle_In : TRANSPORT) return INTEGER is
begin
   return Vehicle_In.Wheels;
end Get_Wheels;

function Get_Weight(Vehicle_In : TRANSPORT) return FLOAT is
begin
   return Vehicle_In.Weight;
end Get_Weight;


-- Subprogram for the CAR record type.
function Tire_Loading(Vehicle_In : CAR) return FLOAT is
begin
   return Vehicle_In.Weight / FLOAT(Vehicle_In.Wheels);
end Tire_Loading;

end Conveyance1;


