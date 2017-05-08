package Conveyance3 is 

   -- This is a very simple transportaion type.
   type TRANSPORT is tagged private;

   procedure Set_Values(Vehicle_In : in out TRANSPORT; 
                        Wheels_In  : INTEGER; 
                        Weight_In  : FLOAT);
   function Get_Wheels(Vehicle_In : TRANSPORT) return INTEGER;
   function Get_Weight(Vehicle_In : TRANSPORT) return FLOAT;


   type CAR is new TRANSPORT with private;

   procedure Set_Values(Vehicle_In : in out CAR; 
                        Passenger_Count_In : INTEGER);
   function Get_Passenger_Count(Vehicle_In : CAR) return INTEGER;


   type TRUCK is new TRANSPORT with private;

   procedure Set_Values(Vehicle_In : in out TRUCK; 
                        Wheels_In : INTEGER; 
                        Weight_In : FLOAT; 
                        Passenger_Count_In : INTEGER; 
                        Payload_In : FLOAT);
   function Get_Passenger_Count(Vehicle_In : TRUCK) return INTEGER;


   type BICYCLE is new TRANSPORT with private;

private

   type TRANSPORT is tagged
      record
         Wheels : INTEGER;
         Weight : FLOAT;
      end record;

   type CAR is new TRANSPORT with
      record
         Passenger_Count : INTEGER;
      end record;

   type TRUCK is new TRANSPORT with
      record
         Passenger_Count : INTEGER;
         Payload         : FLOAT;
      end record;

   type BICYCLE is new TRANSPORT with null record;

end Conveyance3;
