#include "mainwindow.h"
#include <QCoreApplication>
#include <QVBoxLayout>
#include <QLineEdit>

MainWindow::MainWindow(QWidget *parent) : QMainWindow(parent)
{
        // Create the button, make "this" the parent
       m_button = new QPushButton("My Button", this);
       // set size and location of the button
       m_button->setGeometry(QRect(QPoint(100, 100),
       QSize(200, 50)));

       // Connect button signal to appropriate slot
       connect(m_button, SIGNAL (released()), this, SLOT (handleButton()));

       QVBoxLayout *layout = new QVBoxLayout;

       QLabel *title = new QLabel("Guess the world");
       title->setAlignment(Qt::AlignHCenter);
       title->setStyleSheet("QLabel { font-weight : bold; font-size: 25px ; }");
       layout->addWidget(title);

       QHBoxLayout *row1 = new QHBoxLayout;
       QHBoxLayout *row2 = new QHBoxLayout;
       QHBoxLayout *row3 = new QHBoxLayout;

       layout->addItem(row1);
       layout->addItem(row2);
       layout->addItem(row3);

       QLabel *word = new QLabel("Slowo do zgadniecia");
       QLineEdit *blank = new QLineEdit("Insert:");

       row1->addWidget(word);
       row1->addWidget(blank);
       row1->setSpacing(5);

}

void MainWindow::handleButton()
{
    // change the text
    m_button->setText("Example");
    // resize button
    m_button->resize(100,100);
}
