/* CORRECT FORMATTED DEFINITIONS FILE */
#include <defs_for_defs.h>


typedef struct
{
  REC_NO                CorrectDefinitionsChainStart;   /* if TM on Tms e.g pallet of totes */


/* PLUGINSTART (correct_definitions.CORRECT_DEFINITIONS_MEM_REC.plugin.inc)   */
/*pi*/    /* content from file: op.plugin... */
/*pi*/
/*pi*/    REC_NO                XxxNextTmRecNo;
/*pi*/
/* PLUGINEND - end of plugin - edit keyline do not alter                    */

  /* -== ADDER MEM START ==- */

  BOOLEAN Something; //TestComment
  /* -== ADDER MEM END   ==- */

  REC_NO Yyy_ZzzLocNextTmRecNo;

} CORRECT_DEFINITIONS_MEM_REC_TYPE;

typedef struct
{
  ID_TYPE                    TmId;                   /* INDEX */
  CHAIN_TYPE               SelectedDestination;    /* (see above) */
  CHAIN_TYPE               RoutingDestination;     /* (see above) */
  REASON_TYPE              MoveReason;
  int                           ReasonRecNo;            /* needs to be int not REC_NO */

/* PLUGINSTART (tm_definitions.TM_REC_TYPE.plugin.inc)   */
/*pi*/    /* content from file: sorter_chute_loc.plugin... */
/*pi*/
/*pi*/    REC_NO   SorterChuteLocRecNo;
/*pi*/
/* PLUGINEND - end of plugin - edit keyline do not alter                    */


  REASON_TYPE ForceDespatchReqd;

  /* C2066 */
  REC_NO                        LastKnown_UserRecNo;

  /* Used in RTV / STO orders */
  TIME_TYPE                 TimePalletLabelFirstPrinted;

  /* -== ADDER REC START ==- */


  BOOLEAN Alberto; // test
  BOOLEAN Flag1; // hi
  BOOLEAN Flag1; // asdaksjhndkasbd
  /* -== ADDER REC END   ==- */

  /* -== ADDER NONDB START ==- */

  /* DEFNONDBFLD int MyBool; */ // TESTCOMMENT
  /* DEFNONDBFLD int MyBool; */ // bleh
  /* -== ADDERNONDB END   ==- */

  /* DEFNONDBFLD int                       RetryButton; */

} CORRECT_DEFINITIONS_REC_TYPE;