import java.net.*;
import java.util.*;

public class MyReceiver {

    static final int datagramSize=50;
    InetAddress localHost;
    int destinationPort;
    DatagramSocket socket;

    ReceiverThread receiver;

    public MyReceiver(int myPort, int destPort) throws Exception {
        localHost=InetAddress.getByName("127.0.0.1");
        destinationPort = destPort;
        socket = new DatagramSocket(myPort);
        receiver = new ReceiverThread();
        
    }

    public static void main(String[] args) throws Exception {
        MyReceiver receiver = new MyReceiver( Integer.parseInt(args[0]),
                       Integer.parseInt(args[1]));
        receiver.receiver.start();
    }

   class ReceiverThread extends Thread {
        // Na poczatku oczekujemy id
        int expected = 0;
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
                            System.out.println("Added to queue "+ id);
                        }
                        packet.setPort(destinationPort);
                        socket.send(packet);
                    } else {
                        System.out.println("R:" + p.getIntAt(0) + ": " + (char) p.data[4]);
                        packet.setPort(destinationPort);
                        socket.send(packet);
                        expected++;
                        while(queue.containsKey(expected)) {
                            packet = queue.get(expected);
                            queue.remove(expected);
                            p = new Z2Packet(packet.getData());
                            System.out.println("R:" + p.getIntAt(0) + ": " + (char) p.data[4]);
                            expected++;
                            packet.setPort(destinationPort);
                            socket.send(packet);
                        }
                    }
        		}
        	}
            catch(Exception e)
        	{
                System.out.println("Z2Receiver.ReceiverThread.run: "+e);
        	}
        }

    }

}