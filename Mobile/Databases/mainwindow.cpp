#include "mainwindow.h"
#include "ui_mainwindow.h"
#include <QDebug>
#include <QSqlDatabase>
#include <QSqlQuery>
#include <QSqlRecord>
#include <QSqlError>

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    qDebug() << openDB();
    qDebug() << db.lastError();

    QObject::connect(iw,SIGNAL(plsAddStud(QString)),this,SLOT(addStudent(QString data)));
}

MainWindow::~MainWindow()
{
    delete ui;
}

bool MainWindow::openDB(){

    QSqlDatabase::removeDatabase("qt_sql_default_connection");
    db = QSqlDatabase::addDatabase("QSQLITE");
    db.setDatabaseName("my.db.sql");
    return db.open();
}

void MainWindow::on_showB_clicked()
{
    //QString data = "XXX";
    ui->textEdit->clear();
    if(db.isOpen()){
        qDebug() << "dbIsOpen";
        QSqlQuery query("SELECT * FROM students");
        int fieldNo = query.record().indexOf("name");
        while(query.next()){
            ui->textEdit->append(query.value(fieldNo).toString());
        }
        qDebug() << query.lastError();
    }
}

void MainWindow::on_insertB_clicked()
{
    iw = new InsertWindow(this);
    iw->show();
}

void MainWindow::addStudent(QString data){

    if(db.isOpen()){
        qDebug() << data;
        QSqlQuery query("SELECT max(id) FROM students");
        int id = query.value(0).toInt();
        QStringList more = data.split(" ");
        query.prepare("INSERT INTO students values(:id, :name, :surname)");
        query.bindValue(":id",id+1);
        query.bindValue(":name",more.at(0));
        query.bindValue(":surname",more.at(1));
        query.exec();

    } else {
        qDebug() << db.lastError();
    }
}
