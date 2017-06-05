#include <QObject>
#include <QApplication>
#include <QLabel>
#include <QWidget>
#include <QPushButton>
#include <QFormLayout>
#include <QVBoxLayout>
#include <QHBoxLayout>
#include <QLineEdit>
#include "mainwindow.h"

int main(int argc, char* argv[])
{
//    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);
//QGuiApplication app(argc, argv);

//    QApplication app(argc, argv);

//    QWidget *window = new QWidget;

//    QVBoxLayout *layout = new QVBoxLayout;

//    QLabel *title = new QLabel("Guess the world");
//    title->setAlignment(Qt::AlignHCenter);
//    title->setStyleSheet("QLabel { font-weight : bold; font-size: 25px ; }");
//    layout->addWidget(title);

//    QHBoxLayout *row1 = new QHBoxLayout;
//    QHBoxLayout *row2 = new QHBoxLayout;
//    QHBoxLayout *row3 = new QHBoxLayout;

//    layout->addItem(row1);
//    layout->addItem(row2);
//    layout->addItem(row3);

//    QLabel *word = new QLabel("Slowo do zgadniecia");
//    QLineEdit *blank = new QLineEdit("Insert:");

//    row1->addWidget(word);
//    row1->addWidget(blank);
//    row1->setSpacing(5);

//    QLabel *correct = new QLabel("Correct: 0");
//    correct->setStyleSheet("QLabel {color : green; }");
//    correct->setAlignment(Qt::AlignHCenter);

//    QLabel *failed = new QLabel("Failed: 0");
//    failed->setStyleSheet("QLabel {color : red; }");
//    failed->setAlignment(Qt::AlignHCenter);

//    row2->addWidget(correct);
//    row2->addWidget(failed);
//    row2->setSpacing(5);

//    QPushButton *clearB = new QPushButton("CLEAR");
//    QPushButton *nextB = new QPushButton("NEXT");
//    QPushButton *checkB = new QPushButton("CHECK");

//    row3->addWidget(clearB);
//    row3->addWidget(nextB);
//    row3->addWidget(checkB);

//    window->setLayout(layout);
//    window->show();

//    QObject::connect(nextB,SIGNAL(clicked(bool)), word, SLOT(setText(QString)));

//    QQmlApplicationEngine engine;
//    engine.load(QUrl(QLatin1String("qrc:/main.qml")));
//    if (engine.rootObjects().isEmpty())
 //       return -1;

    QApplication app(argc, argv);
    MainWindow mainWindow;
    mainWindow.show();

    return app.exec();
}
