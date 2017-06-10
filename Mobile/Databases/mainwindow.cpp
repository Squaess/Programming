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
//    QSqlQuery l("DROP TABLE students; DROP TABLE courses; DROP TABLE grades");
    QSqlQuery query("CREATE table students(id integer, name varchar(20), surname varchar(20))");
    QSqlQuery q2("CREATE UNIQUE INDEX id on students(id)");
//    QSqlQuery q21("DELETE FROM students");
//    QSqlQuery q3("INSERT INTO students values(1,'Mike','Wazowski')");
//    QSqlQuery q33("INSERT INTO students values(2,'Vincent','Vega')");

    QSqlQuery q4("CREATE TABLE courses(curid integer, curname varchar(20))");
//    QSqlQuery q41("DELETE FROM courses");
    QSqlQuery q45("CREATE UNIQUE INDEX curid on courses(curid)");
//    QSqlQuery q46("INSERT INTO courses values(1,'Aplikacje_Mobilne')");

    QSqlQuery q5("CREATE TABLE grades(studid integer,curid integer, grade integer)");
//    QSqlQuery q6("DELETE FROM grades where studid=1 OR studid=2");
//    QSqlQuery q7("INSERT INTO grades values(1,1,3)");
//    QSqlQuery q8("INSERT INTO grades values(2,1,4)");

//    l.exec();
//    q21.exec();
//    q41.exec();
//    qDebug() << l.lastError();
    query.exec();
    q2.exec();
//    q3.exec();
//    q33.exec();
    q4.exec();
    q45.exec();
//    q46.exec();
    q5.exec();
//    q6.exec();
//    q7.exec();
//    q8.exec();

    qDebug() << db.lastError();
    qDebug() << query.lastError();
    qDebug() << q2.lastError();
//    qDebug() << q3.lastError();
    qDebug() << q4.lastError();
    qDebug() << q45.lastError();
//    qDebug() << q46.lastError();
    qDebug() << q5.lastError();
//    qDebug() << q6.lastError();
//    qDebug() << q7.lastError();


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
    QString data  = "";
    if(db.isOpen()){
        qDebug() << "dbIsOpen";
        QSqlQuery query("SELECT name, surname, curname, grade FROM students inner join grades on id=studid "
                        "inner join courses on grades.curid=courses.curid "
                        "order by name ASC");
        int fieldNo = query.record().indexOf("name");
        int surNo = query.record().indexOf("surname");
        int cN = query.record().indexOf("curname");
        int gN = query.record().indexOf("grade");

        while(query.next()){

            data += query.value(fieldNo).toString()+" ";
            data += query.value(surNo).toString()+ " ";
            data += query.value(cN).toString()+" ";
            data += query.value(gN).toString()+"\n";

            qDebug() << data;
        }
        ui->textEdit->setText(data);
        QSqlQuery queryy("SELECT COUNT(*) FROM grades");
        queryy.first();
        QString data = queryy.record().value(0).toString();
        qDebug() << data;
        QSqlQuery query2("SELECT COUNT(*) FROM students");
        query2.first();
        data = query2.record().value(0).toString();
        qDebug() << data;
        QSqlQuery query3("SELECT COUNT(*) FROM courses");
        query3.first();
        data = query3.record().value(0).toString();
        qDebug() << data;
        qDebug() << "SHOW BUTTON";
    }
}

void MainWindow::on_insertB_clicked()
{
    iw = new InsertWindow(this);
    QObject::connect(iw, SIGNAL(plsAddStud(QString)), this ,SLOT(addStudent(QString)));
    QObject::connect(iw,SIGNAL(plsAddGrade(QString)), this, SLOT(addGrade(QString)));
    QObject::connect(iw,SIGNAL(plsAddCourse(QString)),this, SLOT(addCourse(QString)));
    qDebug() << "insertButtonClicked";
    iw->show();
}

void MainWindow::addStudent(QString data){

    if(db.isOpen()){
        qDebug() << data;
        QSqlQuery query("SELECT max(id) FROM students");
        query.first();
        int id = query.value(0).toInt();
        qDebug() << id;
        QStringList more = data.split(" ");
        qDebug() << more.at(0);
        qDebug() << more.at(1);
        query.prepare("INSERT INTO students values(:id, :name, :surname)");
        query.bindValue(":id",(id+1));
        query.bindValue(":name",more.at(0));
        query.bindValue(":surname",more.at(1));
        query.exec();
        qDebug() << query.lastError();

    } else {
        qDebug() << db.lastError();
    }
}

void MainWindow::addGrade(QString data){
    if(db.isOpen()){
        QStringList more = data.split(" ");
        QSqlQuery quer;
        quer.prepare("INSERT INTO grades values(:sid, :cid, :grade)");
        quer.bindValue(":sid",more.at(0));
        quer.bindValue(":cid",more.at(1));
        quer.bindValue(":grade", more.at(2));
        quer.exec();
    }
}

void MainWindow::addCourse(QString data){
    if(db.isOpen()){
        QSqlQuery query("SELECT max(curid) FROM courses");
        query.first();
        int id = query.value(0).toInt();

        query.prepare("INSERT INTO courses values(:id, :name)");
        query.bindValue(":id",(id+1));
        query.bindValue(":name",data);
        query.exec();
    }
}

void MainWindow::on_ssB_clicked()
{
    QString data = "";
    if(db.isOpen()){
        qDebug() << "dbIsOpen";
        QSqlQuery query("SELECT * FROM students");
        int id = query.record().indexOf("id");
        int nameNo = query.record().indexOf("name");
        int surNo = query.record().indexOf("surname");

        while(query.next()){
            data += query.value(id).toString()+" ";
            data += query.value(nameNo).toString()+" ";
            data += query.value(surNo).toString()+ "\n";

            qDebug() << data;
        }
        ui->textEdit->setText(data);
    }
}

void MainWindow::on_scB_clicked()
{
    QString data = "";
    if(db.isOpen()){
        qDebug() << "dbIsOpen";
        QSqlQuery query("SELECT * FROM courses");
        int id = query.record().indexOf("curid");
        int nameNo = query.record().indexOf("curname");

        while(query.next()){
            data += query.value(id).toString()+" ";
            data += query.value(nameNo).toString()+"\n";

            qDebug() << data;
        }
        ui->textEdit->setText(data);
    }
}

void MainWindow::on_deleteB_clicked()
{
    dw = new DeleteWindow(this);
    QObject::connect(dw,SIGNAL(plsDeleteStudent(QString)),this, SLOT(deleteStudent(QString)));
    QObject::connect(dw, SIGNAL(plsDeleteCourse(QString)),this, SLOT(deleteCourse(QString)));
    dw->show();
}

void MainWindow::deleteStudent(QString data){
    if(db.isOpen()){
        qDebug() << "dbIsOpen";
        QSqlQuery query;
        query.prepare("DELETE FROM students where id = :idd");
        query.bindValue(":idd",data);
        query.exec();
        qDebug() << query.lastError();
        query.prepare("DELETE FROM grades where studid=:id");
        query.bindValue(":id",data);
        query.exec();
        qDebug() << query.lastError();
    }
}

void MainWindow::deleteCourse(QString data){
    if (db.isOpen()){
        qDebug() << "dbisOpen";
        QSqlQuery query;
        query.prepare("DELETE FROM courses where curid = :id");
        query.bindValue(":id",data);
        query.exec();
        query.prepare("DELETE FROM grades WHERE curid = :id");
        query.bindValue(":id",data);
        query.exec();
    }
}

void MainWindow::on_searchB_clicked()
{
    sw = new SearchWindow(this);
    QObject::connect(sw,SIGNAL(plsSearchStudName(QString)),this,SLOT(searchStudName(QString)));
    QObject::connect(sw,SIGNAL(plsSearchStudName(QString)),sw, SLOT(close()));
    QObject::connect(sw, SIGNAL(plsSearchStudSur(QString)),this,SLOT(searchStudSur(QString)));
    QObject::connect(sw,SIGNAL(plsSearchStudSur(QString)),sw,SLOT(close()));
    QObject::connect(sw, SIGNAL(plsSearchCours(QString)),this,SLOT(searchCourse(QString)));
    QObject::connect(sw,SIGNAL(plsSearchCours(QString)),sw,SLOT(close()));
    sw->show();
}

void MainWindow::searchStudName(QString data){

    QString out = "";
    if(db.isOpen()){
        qDebug() << "dbIsOpen";
        qDebug() << data;
        QSqlQuery query;
        query.prepare("SELECT * FROM students WHERE name = :sn");
        query.bindValue(":sn",data);
        query.exec();
        int id = query.record().indexOf("id");
        int nameNo = query.record().indexOf("name");
        int surNo = query.record().indexOf("surname");

        while(query.next()){
            out += query.value(id).toString()+" ";
            out += query.value(nameNo).toString()+" ";
            out += query.value(surNo).toString()+ "\n";

            qDebug() << out;
        }
        ui->textEdit->setText(out);
    }
}

void MainWindow::searchStudSur(QString data){
    QString out = "";
    if(db.isOpen()){
        qDebug() << "dbIsOpen";
        qDebug() << data;
        QSqlQuery query;
        query.prepare("SELECT * FROM students WHERE surname = :sn");
        query.bindValue(":sn",data);
        query.exec();
        int id = query.record().indexOf("id");
        int nameNo = query.record().indexOf("name");
        int surNo = query.record().indexOf("surname");

        while(query.next()){
            out += query.value(id).toString()+" ";
            out += query.value(nameNo).toString()+" ";
            out += query.value(surNo).toString()+ "\n";

            qDebug() << out;
        }
        ui->textEdit->setText(out);
    }
}

void MainWindow::searchCourse(QString data){
    QString out = "";
    if(db.isOpen()){
        qDebug() << "dbIsOpen";
        qDebug() << data;
        QSqlQuery query;
        query.prepare("SELECT * FROM courses WHERE curname = :sn");
        query.bindValue(":sn",data);
        query.exec();
        int id = query.record().indexOf("curid");
        int nameNo = query.record().indexOf("curname");

        while(query.next()){
            out += query.value(id).toString()+" ";
            out += query.value(nameNo).toString()+"\n";

            qDebug() << out;
        }
        ui->textEdit->setText(out);
    }
}
