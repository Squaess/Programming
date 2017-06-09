import java.net.*;
import java.util.*;

public class Receiver {

    private static final int datagramSize = 50;
    private InetAddress localHost;
    private int destinationPort;
    private DatagramSocket socket;
    // CZAS PO JAKIM PONAWIANA JEST PROSBA O NASTEPNY PAKIET
    private int sleepTime = 7000;

    private ReceiverThread receiver;
    private ResendThread resender;

    private int expected = 0;

    private Receiver(int myPort, int destPort) throws Exception {
        localHost=InetAddress.getByName("127.0.0.1");
        destinationPort = destPort;
        socket = new DatagramSocket(myPort);
        receiver = new ReceiverThread();
        resender = new ResendThread();
    }

    public static void main(String[] args) throws Exception {
        Receiver receiver = new Receiver( Integer.parseInt(args[0]),
                       Integer.parseInt(args[1]));
        receiver.receiver.start();
        receiver.resender.start();
    }

   class ReceiverThread extends Thread {
        // Na poczatku oczekujemy id
        
        Hashtable<Integer, DatagramPacket> queue;

        public void run() {
            queue = new Hashtable<Integer, DatagramPacket>();
            try {
        	    while(true)
        		{
        		    byte[] data = new byte[datagramSize];
        		    DatagramPacket packet = new DatagramPacket(data, datagramSize);
        		    socket.receive(packet);

                    Z2Packet p = new Z2Packet(packet.getData());
                    int id = p.getIntAt(0);

                    if(id != expected) {

                        if(!queue.containsKey(id) && id >= expected) {
                            queue.put(id, packet);
                            System.out.println("queue "+ id);
                        }

                    } else {
                        // Pdebralismy pakiet na ktory czekamy
                        System.out.println("R:" + p.getIntAt(0) + ": " + (char) p.data[4]);
                        expected++;
                        // sprawdzamy czy wczesniej nie odebralismy juz nastepnych pakietow
                        while(queue.containsKey(expected)) {
                            packet = queue.get(expected);
                            queue.remove(expected);
                            p = new Z2Packet(packet.getData());
                            System.out.println("R:" + p.getIntAt(0) + ": " + (char) p.data[4]);
                            expected++;
                        }
                        // wysylamy ostatni pakiet jaki mamy wyswietlony
                        packet.setPort(destinationPort);
                        socket.send(packet);
                        System.out.println("next: "+ expected);
                    }
        		}
        	}
            catch(Exception e)
        	{
                System.out.println("Z2Receiver.ReceiverThread.run: "+e);
        	}
        }

    }

    class ResendThread extends Thread {
        int waiting = 0;
        public void run() {
            try {
                while(true){
                    sleep(sleepTime);
                    if(waiting == expected){
                        Z2Packet p = new Z2Packet(4+1);
                        p.setIntAt(waiting-1,0);
                        p.data[4] = (byte) 'R';
                        DatagramPacket packet = new DatagramPacket(p.data, p.data.length, localHost, destinationPort);
                        socket.send(packet);
                        System.out.println("Ponawiam prosbe "+ waiting);
                    } else {
                        waiting = expected;
                    }
                }
            } catch (Exception e){
                e.printStackTrace();
            }
        }
    }

}