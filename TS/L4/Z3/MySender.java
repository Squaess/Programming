import java.net.*;
import java.util.*;


class MySender {

	static final int datagramSize=50;
	static final int sleepTime=500;
	static final int maxPacket=50;

	static final int retransmisisonTime=1000;

	InetAddress localHost;
	int destinationPort;
	DatagramSocket socket;

	SenderThread sender;
	ReceiverThread receiver;

	ReSenderThread resender;
	int sentPackets = 0;
	int receivedPackets = 0;
	Hashtable<Integer, DatagramPacket> queue;
	ArrayList<Byte> orgData = new ArrayList<Byte>();

	public MySender(int myPort, int destPort) throws Exception {
	    localHost=InetAddress.getByName("127.0.0.1");
	    destinationPort = destPort;
	    socket = new DatagramSocket(myPort);
	    sender = new SenderThread();
	    receiver = new ReceiverThread();
	    resender = new ReSenderThread();
	    queue = new Hashtable<Integer, DatagramPacket>();
	}

	public static void main(String[] args) throws Exception {

		MySender sender=new MySender( Integer.parseInt(args[0]),
					   Integer.parseInt(args[1]));
		sender.sender.start();
	    sender.receiver.start();
	    sender.resender.start();
	}

	class SenderThread extends Thread {
		public void run() {
		    int i, x;
		    try {
			    for(i=0; (x=System.in.read()) >= 0 ; i++) {
			    	orgData.add((byte) x);
		            Z2Packet p = new Z2Packet(4+1);
		            p.setIntAt(i,0);
		            p.data[4]= (byte) x;
				    DatagramPacket packet = 
					new DatagramPacket(p.data, p.data.length, 
							   localHost, destinationPort);
				    socket.send(packet);
				    System.out.println("Send: "+ i);
					sentPackets++;
				    sleep(sleepTime);
				}
			}
		    catch(Exception e)
			{
		        System.out.println("Z2Sender.SenderThread.run: "+e);
			}
		}
	}

	class ReceiverThread extends Thread {

		public void run() {
		    try
			{
			    while(true)
				{
				    byte[] data = new byte[datagramSize];
				    DatagramPacket packet =	new DatagramPacket(data, datagramSize);
				    socket.receive(packet);
		            
		            Z2Packet p = new Z2Packet(packet.getData());
				    System.out.println("Odebralem:"+p.getIntAt(0)+
												": "+ (char)p.data[4]);		
		            int id = p.getIntAt(0);
					if (id != receivedPackets) {
						if(!queue.containsKey(id) && id >= receivedPackets)	{
							queue.put(id, packet);
							System.out.printf("-q: %d\n", id);
						}
					} else {
						System.out.println("S:" + p.getIntAt(0) + ": " + (char) p.data[4]);
						receivedPackets++;
							while (queue.containsKey(receivedPackets)) {
								packet = queue.get(receivedPackets);
							//	queue.remove(receivedPackets);
								p = new Z2Packet(packet.getData());
								System.out.println("s:" + p.getIntAt(0) + ": " + (char) p.data[4]);
								receivedPackets++;
							}
					}
				}
				
			}
		    catch(Exception e)
			{
		        System.out.println("Z2Sender.ReceiverThread.run: ");
		        e.printStackTrace();
			}
		}

	}

	class ReSenderThread extends Thread {
		public void run() {
			int lastPacket = 0;
			try {
				while (true) {
					sleep(sleepTime * 5);

					if (lastPacket == receivedPackets) {

						for (int i = receivedPackets; i < sentPackets; i++) {
							if (queue.containsKey (i))
								continue;
							Z2Packet p = new Z2Packet (4 + 1);
							p.setIntAt (receivedPackets, 0);
							p.data[4] = (byte) orgData.get(receivedPackets);
							DatagramPacket packet =
							new DatagramPacket(p.data, p.data.length,
							localHost, destinationPort);
							socket.send(packet);
							System.out.println("Resend: "+ i);
							sleep(sleepTime);
						}
					} else
						lastPacket = receivedPackets;
				}
			} catch (Exception e) {
				System.out.println("Z2Sender.SenderThread.run: " + e);
			}
		}
	}
}